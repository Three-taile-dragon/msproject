package project_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"test.com/project_common/encrypts"
	"test.com/project_common/errs"
	"test.com/project_grpc/project"
	"test.com/project_project/config"
	"test.com/project_project/internal/dao"
	"test.com/project_project/internal/dao/mysql"
	"test.com/project_project/internal/data/menu"
	"test.com/project_project/internal/database/tran"
	"test.com/project_project/internal/repo"
	"test.com/project_project/pkg/model"
)

// ProjectService grpc 登陆服务 实现
type ProjectService struct {
	project.UnimplementedProjectServiceServer
	cache       repo.Cache
	transaction tran.Transaction
	menuRepo    repo.MenuRepo
	projectRepo repo.ProjectRepo
}

func New() *ProjectService {
	return &ProjectService{
		cache:       dao.Rc,
		transaction: dao.NewTransaction(),
		menuRepo:    mysql.NewMenuDao(),
		projectRepo: mysql.NewProjectDao(),
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
	pms, total, err := p.projectRepo.FindProjectByMemId(ctx, memberId, page, pageSize)
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
		v.Code, _ = encrypts.EncryptInt64(v.Id, config.C.AC.AesKey)
	}
	return &project.MyProjectResponse{Pm: pmm, Total: total}, nil
}
