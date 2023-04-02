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
	"test.com/project_project/internal/dao"
	"test.com/project_project/internal/dao/mysql"
	"test.com/project_project/internal/data/menu"
	pro "test.com/project_project/internal/data/project"
	"test.com/project_project/internal/data/task"
	"test.com/project_project/internal/database/tran"
	"test.com/project_project/internal/repo"
	"test.com/project_project/pkg/model"
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
func (p *ProjectService) Index(ctx context.Context, req *project.IndexRequest) (*project.IndexResponse, error) {
	c := context.Background()
	pms, err := p.menuRepo.FindMenus(c)
	if err != nil {
		zap.L().Error("首页模块menu数据库存入出错", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	childs := menu.CovertChild(pms)
	var mms []*project.MenuMessage
	err = copier.Copy(&mms, childs)
	if err != nil {
		zap.L().Error("首页模块childs结构体赋值错误", zap.Error(err))
		return nil, errs.GrpcError(model.CopyError)
	}
	return &project.IndexResponse{Menus: mms}, nil
}

func (p *ProjectService) FindProjectByMemId(ctx context.Context, req *project.ProjectRpcMessage) (*project.MyProjectResponse, error) {
	memberId := req.MemberId
	page := req.Page
	pageSize := req.PageSize
	var pms []*pro.ProjectAndMember
	var total int64
	var err error
	if req.SelectBy == "" || req.SelectBy == "my" {
		pms, total, err = p.projectRepo.FindProjectByMemId(ctx, memberId, "", page, pageSize)
	}
	if req.SelectBy == "archive" {
		pms, total, err = p.projectRepo.FindProjectByMemId(ctx, memberId, "and archive = 1", page, pageSize)
	}
	if req.SelectBy == "deleted" {
		pms, total, err = p.projectRepo.FindProjectByMemId(ctx, memberId, "and deleted = 1", page, pageSize)
	}
	if req.SelectBy == "collect" {
		//跨表查询
		pms, total, err = p.projectRepo.FindCollectProjectByMemId(ctx, memberId, page, pageSize)
	}

	if err != nil {
		zap.L().Error("首页模块project查找失败", zap.Error(err))
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
		zap.L().Error("首页模块pmm结构体赋值错误", zap.Error(err))
		return nil, errs.GrpcError(model.CopyError)
	}
	for _, v := range pmm {
		v.Code, _ = encrypts.EncryptInt64(v.Id, model.AesKey)
		//v.Code = strconv.FormatInt(v.Id, 10)
		pam := pro.ToMap(pms)[v.Id]
		v.AccessControlType = pam.GetAccessControlType()
		v.OrganizationCode, _ = encrypts.EncryptInt64(pam.OrganizationCode, model.AesKey)
		//v.OrganizationCode = strconv.FormatInt(pam.OrganizationCode, 10)
		v.JoinTime = tms.FormatByMill(pam.JoinTime)
		v.OwnerName = req.MemberName
		v.Order = int32(pam.Sort)
		v.CreateTime = tms.FormatByMill(pam.CreateTime)
	}
	return &project.MyProjectResponse{Pm: pmm, Total: total}, nil
}

func (p *ProjectService) FindProjectTemplate(ctx context.Context, req *project.ProjectRpcMessage) (*project.ProjectTemplateResponse, error) {
	//1. 根据viewType去查询项目模板表 得到list
	organizationCodeStr, _ := encrypts.Decrypt(req.OrganizationCode, model.AesKey) //解密操作
	organizationCode, _ := strconv.ParseInt(organizationCodeStr, 10, 64)
	page := req.Page
	pageSize := req.PageSize
	//ViewType 	-1 代表查询全部模板
	var pts []pro.ProjectTemplate
	var total int64
	var err error
	if req.ViewType == -1 {
		pts, total, err = p.projectTemplateRepo.FindProjectTemplateAll(ctx, organizationCode, page, pageSize)
	}
	//ViewType 	0 代表查询自定义模板
	if req.ViewType == 0 {
		pts, total, err = p.projectTemplateRepo.FindProjectTemplateCustom(ctx, req.MemberId, organizationCode, page, pageSize)
	}
	//ViewType 	1 代表查询系统模板
	if req.ViewType == 1 {
		pts, total, err = p.projectTemplateRepo.FindProjectTemplateSystem(ctx, page, pageSize)
	}
	if err != nil {
		zap.L().Error("项目模板模块类型选择出错", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	//2. 模型转换 拿到模板id列表，去 任务步骤模板表，去进行查询
	tsts, err := p.taskStagesTemplateRepo.FindInProTemIds(ctx, pro.ToProjectTemplateIds(pts))
	if err != nil {
		zap.L().Error("项目模板模型转换", zap.Error(err))
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
		zap.L().Error("项目模块赋值错误", zap.Error(err))
		return nil, errs.GrpcError(model.CopyError)
	}
	return &project.ProjectTemplateResponse{Ptm: pmMsgs, Total: total}, nil
}
