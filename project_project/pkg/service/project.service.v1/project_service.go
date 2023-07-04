package project_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"strconv"
	"test.com/project_common/encrypts"
	"test.com/project_common/errs"
	"test.com/project_common/tms"
	"test.com/project_grpc/project"
	"test.com/project_grpc/user/login"
	"test.com/project_project/internal/dao"
	"test.com/project_project/internal/dao/mysql"
	"test.com/project_project/internal/data/menu"
	pro "test.com/project_project/internal/data/project"
	"test.com/project_project/internal/data/task"
	"test.com/project_project/internal/database"
	"test.com/project_project/internal/database/tran"
	"test.com/project_project/internal/repo"
	"test.com/project_project/internal/rpc"
	"test.com/project_project/pkg/model"
	"time"
)

// ProjectService grpc 登陆服务 实现
type ProjectService struct {
	project.UnimplementedProjectServiceServer
	cache                  repo.Cache
	transaction            tran.Transaction
	menuRepo               repo.MenuRepo
	projectRepo            repo.ProjectRepo
	projectTemplateRepo    repo.ProjectTemplateRepo
	taskStagesTemplateRepo repo.TaskStagesTemplateRepo
}

func New() *ProjectService {
	return &ProjectService{
		cache:                  dao.Rc,
		transaction:            dao.NewTransaction(),
		menuRepo:               mysql.NewMenuDao(),
		projectRepo:            mysql.NewProjectDao(),
		projectTemplateRepo:    mysql.NewProjectTemplateDao(),
		taskStagesTemplateRepo: mysql.NewTaskStagesTemplateDao(),
	}
}
func (ps *ProjectService) Index(ctx context.Context, req *project.IndexRequest) (*project.IndexResponse, error) {
	c := context.Background()
	pms, err := ps.menuRepo.FindMenus(c)
	if err != nil {
		zap.L().Error("project Index FindMenus error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	childs := menu.CovertChild(pms)
	var mms []*project.MenuMessage
	err = copier.Copy(&mms, childs)
	if err != nil {
		zap.L().Error("project Index Copy error", zap.Error(err))
		return nil, errs.GrpcError(model.CopyError)
	}
	return &project.IndexResponse{Menus: mms}, nil
}

func (ps *ProjectService) FindProjectByMemId(ctx context.Context, req *project.ProjectRpcMessage) (*project.MyProjectResponse, error) {
	memberId := req.MemberId
	page := req.Page
	pageSize := req.PageSize
	var pms []*pro.ProjectAndMember
	var total int64
	var err error
	if req.SelectBy == "" || req.SelectBy == "my" {
		pms, total, err = ps.projectRepo.FindProjectByMemId(ctx, memberId, "and deleted = 0", page, pageSize)
	}
	if req.SelectBy == "archive" {
		pms, total, err = ps.projectRepo.FindProjectByMemId(ctx, memberId, "and archive = 1", page, pageSize)
	}
	if req.SelectBy == "deleted" {
		pms, total, err = ps.projectRepo.FindProjectByMemId(ctx, memberId, "and deleted = 1", page, pageSize)
	}
	if req.SelectBy == "collect" {
		//跨表查询
		pms, total, err = ps.projectRepo.FindCollectProjectByMemId(ctx, memberId, page, pageSize)
	}

	if err != nil {
		zap.L().Error("project FindProjectByMemId FindProjectByMemId/FindCollectProjectByMemId error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	//如果查询的项目数量为空 则返回空值
	if pms == nil {
		return &project.MyProjectResponse{Pm: []*project.ProjectMessage{}, Total: total}, nil
	}
	//拷贝数据
	var pmm []*project.ProjectMessage
	err = copier.Copy(&pmm, pms)
	if err != nil {
		zap.L().Error("project FindProjectByMemId Copy error", zap.Error(err))
		return nil, errs.GrpcError(model.CopyError)
	}
	for _, v := range pmm {
		v.Code, _ = encrypts.EncryptInt64(v.Id, model.AESKey)
		//v.Code = strconv.FormatInt(v.Id, 10)
		pam := pro.ToMap(pms)[v.Id]
		v.AccessControlType = pam.GetAccessControlType()
		v.OrganizationCode, _ = encrypts.EncryptInt64(pam.OrganizationCode, model.AESKey)
		//v.OrganizationCode = strconv.FormatInt(pam.OrganizationCode, 10)
		v.JoinTime = tms.FormatByMill(pam.JoinTime)
		v.OwnerName = req.MemberName
		v.Order = int32(pam.Sort)
		v.CreateTime = tms.FormatByMill(pam.CreateTime)
	}
	return &project.MyProjectResponse{Pm: pmm, Total: total}, nil
}

func (ps *ProjectService) FindProjectTemplate(ctx context.Context, req *project.ProjectRpcMessage) (*project.ProjectTemplateResponse, error) {
	//1. 根据viewType去查询项目模板表 得到list
	organizationCodeStr, _ := encrypts.Decrypt(req.OrganizationCode, model.AESKey) //解密操作
	organizationCode, _ := strconv.ParseInt(organizationCodeStr, 10, 64)
	page := req.Page
	pageSize := req.PageSize
	//ViewType 	-1 代表查询全部模板
	var pts []pro.ProjectTemplate
	var total int64
	var err error
	if req.ViewType == -1 {
		pts, total, err = ps.projectTemplateRepo.FindProjectTemplateAll(ctx, organizationCode, page, pageSize)
	}
	//ViewType 	0 代表查询自定义模板
	if req.ViewType == 0 {
		pts, total, err = ps.projectTemplateRepo.FindProjectTemplateCustom(ctx, req.MemberId, organizationCode, page, pageSize)
	}
	//ViewType 	1 代表查询系统模板
	if req.ViewType == 1 {
		pts, total, err = ps.projectTemplateRepo.FindProjectTemplateSystem(ctx, page, pageSize)
	}
	if err != nil {
		zap.L().Error("project FindProjectTemplate FindProjectTemplateAll/Custom/System error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	//2. 模型转换 拿到模板id列表，去 任务步骤模板表，去进行查询
	tsts, err := ps.taskStagesTemplateRepo.FindInProTemIds(ctx, pro.ToProjectTemplateIds(pts))
	if err != nil {
		zap.L().Error("project FindProjectTemplate FindInProTemIds error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	var ptas []*pro.ProjectTemplateAll
	for _, v := range pts {
		//改谁做的事一定要交出去
		ptas = append(ptas, v.Convert(task.CovertProjectMap(tsts)[v.Id]))
	}
	//3. 组装数据
	var pmMsgs []*project.ProjectTemplate
	err = copier.Copy(&pmMsgs, ptas)
	if err != nil {
		zap.L().Error("project FindProjectTemplate Copy error", zap.Error(err))
		return nil, errs.GrpcError(model.CopyError)
	}
	return &project.ProjectTemplateResponse{Ptm: pmMsgs, Total: total}, nil
}

func (ps *ProjectService) SaveProject(ctx context.Context, msg *project.ProjectRpcMessage) (*project.SaveProjectMessage, error) {
	organizationCodeStr, _ := encrypts.Decrypt(msg.OrganizationCode, model.AESKey)
	organizationCode, _ := strconv.ParseInt(organizationCodeStr, 10, 64)
	templateCodeStr, _ := encrypts.Decrypt(msg.TemplateCode, model.AESKey)
	templateCode, _ := strconv.ParseInt(templateCodeStr, 10, 64)
	//1. 保存项目表
	pr := &pro.Project{
		Name:              msg.Name,
		Description:       msg.Description,
		TemplateCode:      int(templateCode),
		CreateTime:        time.Now().UnixMilli(),
		Cover:             "https://img2.baidu.com/it/u=792555388,2449797505&fm=253&fmt=auto&app=138&f=JPEG?w=667&h=500",
		Deleted:           model.NoDeleted,
		Archive:           model.NoArchive,
		OrganizationCode:  organizationCode,
		AccessControlType: model.Open,
		TaskBoardTheme:    model.Simple,
	}
	//添加事务
	err := ps.transaction.Action(func(conn database.DbConn) error {
		err := ps.projectRepo.SaveProject(conn, ctx, pr)
		if err != nil {
			zap.L().Error("project SaveProject SaveProject error", zap.Error(err))
			return errs.GrpcError(model.DBError)
		}
		pm := &pro.ProjectMember{
			ProjectCode: pr.Id,
			MemberCode:  msg.MemberId,
			JoinTime:    time.Now().UnixMilli(),
			IsOwner:     msg.MemberId,
			Authorize:   "",
		}
		//2. 保存项目和成员的关联表
		err = ps.projectRepo.SaveProjectMember(conn, ctx, pm)
		if err != nil {
			zap.L().Error("project SaveProject SaveProjectMember error", zap.Error(err))
			return errs.GrpcError(model.DBError)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	code, _ := encrypts.EncryptInt64(pr.Id, model.AESKey)
	rsp := &project.SaveProjectMessage{
		Id:               pr.Id,
		Code:             code,
		OrganizationCode: organizationCodeStr,
		Name:             pr.Name,
		Cover:            pr.Cover,
		CreateTime:       tms.FormatByMill(pr.CreateTime),
		TaskBoardTheme:   pr.TaskBoardTheme,
	}

	return rsp, nil
}

// FindProjectDetail 读取项目
func (ps *ProjectService) FindProjectDetail(ctx context.Context, msg *project.ProjectRpcMessage) (*project.ProjectDetailMessage, error) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second) //编写上下文 最多允许两秒超时
	defer cancel()
	cipherIdCodeStr, _ := encrypts.Decrypt(msg.ProjectCode, model.AESKey)
	cipherIdCode, _ := strconv.ParseInt(cipherIdCodeStr, 10, 64)
	// 转换ID 将加密所用的ID转为真正的项目ID

	projectCode, err := ps.projectRepo.FindProjectByCipId(c, cipherIdCode)
	if err != nil {
		zap.L().Error("project FindProjectDetail FindProjectByCipId error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}

	memberId := msg.MemberId
	projectAndMember, err := ps.projectRepo.FindProjectByPIdAndMemId(c, projectCode, memberId)
	if err != nil {
		zap.L().Error("project FindProjectDetail FindProjectByPIdAndMemId error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	ownerId := projectAndMember.IsOwner
	member, err := rpc.LoginServiceClient.FindMemInfoById(c, &login.UserMessage{MemId: ownerId})
	if err != nil {
		zap.L().Error("project rpc FindProjectDetail FindMemInfoById error", zap.Error(err))
		return nil, err
	}
	//TODO  收藏可以放入Redis
	isCollect, err := ps.projectRepo.FindCollectByPIdAndMemId(c, projectCode, memberId)
	if err != nil {
		zap.L().Error("project FindProjectDetail FindCollectByPIdAndMemId error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if isCollect {
		projectAndMember.Collected = model.Collected
	} else {
		projectAndMember.Collected = model.NoCollected
	}
	var detailMsg = &project.ProjectDetailMessage{}
	err = copier.Copy(detailMsg, projectAndMember)
	if err != nil {
		zap.L().Error("project FindProjectDetail Copy error", zap.Error(err))
		return nil, errs.GrpcError(model.CopyError)
	}
	// TODO
	detailMsg.OwnerName = member.Name
	detailMsg.OwnerAvatar = member.Avatar
	detailMsg.Code, _ = encrypts.EncryptInt64(projectAndMember.Id, model.AESKey)
	detailMsg.AccessControlType = projectAndMember.GetAccessControlType()
	detailMsg.OrganizationCode, _ = encrypts.EncryptInt64(projectAndMember.OrganizationCode, model.AESKey)
	detailMsg.Order = int32(projectAndMember.Sort)
	detailMsg.CreateTime = tms.FormatByMill(projectAndMember.CreateTime)
	return detailMsg, err
}

//1. 查项目表
//2. 查项目和成员的关联表 查项目的拥有者 去member表查名字
//3. 查收藏表 判断收藏状态

func (ps *ProjectService) RecycleProject(ctx context.Context, msg *project.ProjectRpcMessage) (*project.RecycleProjectResponse, error) {
	//项目回收
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second) //编写上下文 最多允许两秒超时
	defer cancel()
	cipherIdCodeStr, _ := encrypts.Decrypt(msg.ProjectCode, model.AESKey)
	cipherIdCode, _ := strconv.ParseInt(cipherIdCodeStr, 10, 64)
	projectCode, err := ps.projectRepo.FindProjectByCipId(c, cipherIdCode)
	if err != nil {
		zap.L().Error("project RecycleProject FindProjectByCipId error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	err = ps.projectRepo.DeleteProject(c, projectCode)
	if err != nil {
		zap.L().Error("project RecycleProject DeleteProject error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	return &project.RecycleProjectResponse{}, nil
}

func (ps *ProjectService) RecoveryProject(ctx context.Context, msg *project.ProjectRpcMessage) (*project.RecoveryProjectResponse, error) {
	//从回收站恢复项目
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second) //编写上下文 最多允许两秒超时
	defer cancel()
	cipherIdCodeStr, _ := encrypts.Decrypt(msg.ProjectCode, model.AESKey)
	cipherIdCode, _ := strconv.ParseInt(cipherIdCodeStr, 10, 64)
	projectCode, err := ps.projectRepo.FindProjectByCipId(c, cipherIdCode)
	if err != nil {
		zap.L().Error("project RecycleProject FindProjectByCipId error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	err = ps.projectRepo.RecoveryProject(c, projectCode)
	if err != nil {
		zap.L().Error("project RecycleProject DeleteProject error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	return &project.RecoveryProjectResponse{}, nil
}
