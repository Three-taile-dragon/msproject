package menu_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"test.com/project_common/errs"
	"test.com/project_grpc/menu"
	"test.com/project_project/internal/dao"
	"test.com/project_project/internal/database/tran"
	"test.com/project_project/internal/domain"
	"test.com/project_project/internal/repo"
)

// MenuService grpc 登陆服务 实现
type MenuService struct {
	menu.UnimplementedMenuServiceServer
	cache       repo.Cache
	transaction tran.Transaction
	menuDomain  *domain.MenuDomain
}

func New() *MenuService {
	return &MenuService{
		cache:       dao.Rc,
		transaction: dao.NewTransaction(),
		menuDomain:  domain.NewMenuDomain(),
	}
}
func (m *MenuService) MenuList(context.Context, *menu.MenuReqMessage) (*menu.MenuResponseMessage, error) {
	treeList, err := m.menuDomain.MenuTreeList()
	if err != nil {
		zap.L().Error("project menu MenuList menuDomain.MenuTreeList error", zap.Error(err))
		return nil, errs.GrpcError(err)
	}
	var list []*menu.MenuMessage
	_ = copier.Copy(&list, treeList)
	return &menu.MenuResponseMessage{List: list}, nil
}
