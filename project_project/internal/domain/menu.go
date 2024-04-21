package domain

import (
	"context"
	"go.uber.org/zap"
	"test.com/project_common/errs"
	"test.com/project_project/internal/dao/mysql"
	"test.com/project_project/internal/data/menu"
	"test.com/project_project/internal/repo"
	"test.com/project_project/pkg/model"
)

type MenuDomain struct {
	menuRepo repo.MenuRepo
}

func NewMenuDomain() *MenuDomain {
	return &MenuDomain{
		menuRepo: mysql.NewMenuDao(),
	}
}

func (d *MenuDomain) MenuTreeList() ([]*menu.ProjectMenuChild, *errs.BError) {
	menus, err := d.menuRepo.FindMenus(context.Background())
	if err != nil {
		zap.L().Error("project Menu MenuTreeList menuRepo.FindMenus error", zap.Error(err))
		return nil, model.DBError
	}
	menuChildren := menu.CovertChild(menus)
	return menuChildren, nil
}
