package project_service_v1

import (
	"context"
	"test.com/project_grpc/project"
	"test.com/project_project/internal/dao"
	"test.com/project_project/internal/database/tran"
	"test.com/project_project/internal/repo"
)

// ProjectService LoginService grpc 登陆服务 实现
type ProjectService struct {
	project.UnimplementedProjectServiceServer
	cache       repo.Cache
	transaction tran.Transaction
}

func (p *ProjectService) Index(ctx context.Context, req *project.IndexRequest) (*project.IndexResponse, error) {
	//TODO implement me
	panic("implement me")
}

func New() *ProjectService {
	return &ProjectService{
		cache:       dao.Rc,
		transaction: dao.NewTransaction(),
	}
}
