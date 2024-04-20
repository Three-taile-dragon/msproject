package account_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"test.com/project_common/encrypts"
	"test.com/project_common/errs"
	"test.com/project_grpc/auth"
	"test.com/project_project/internal/dao"
	"test.com/project_project/internal/database/tran"
	"test.com/project_project/internal/domain"
	"test.com/project_project/internal/repo"
)

// AuthService grpc 登陆服务 实现
type AuthService struct {
	auth.UnimplementedAuthServiceServer
	cache             repo.Cache
	transaction       tran.Transaction
	projectAuthDomain *domain.ProjectAuthDomain
}

func New() *AuthService {
	return &AuthService{
		cache:             dao.Rc,
		transaction:       dao.NewTransaction(),
		projectAuthDomain: domain.NewProjectAuthDomain(),
	}
}

func (a *AuthService) AuthList(ctx context.Context, msg *auth.AuthReqMessage) (*auth.ListAuthMessage, error) {
	organizationCode := encrypts.DecryptNoErr(msg.OrganizationCode)
	listPage, total, err := a.projectAuthDomain.AuthListPage(organizationCode, msg.Page, msg.PageSize)
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	var prList []*auth.ProjectAuth
	_ = copier.Copy(&prList, listPage)
	return &auth.ListAuthMessage{List: prList, Total: total}, nil
}
