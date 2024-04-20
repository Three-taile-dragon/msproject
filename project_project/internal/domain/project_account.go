package domain

import (
	"context"
	"go.uber.org/zap"
	"test.com/project_common/errs"
	"test.com/project_project/internal/dao/mysql"
	"test.com/project_project/internal/data/account"
	"test.com/project_project/internal/repo"
	"test.com/project_project/pkg/model"
	"time"
)

type ProjectAuthDomain struct {
	projectAuthRepo repo.ProjectAuthRepo
}

func NewProjectAuthDomain() *ProjectAuthDomain {
	return &ProjectAuthDomain{
		projectAuthRepo: mysql.NewProjectAuthDao(),
	}
}

func (d *ProjectAuthDomain) AuthList(orgCode int64) ([]*account.ProjectAuthDisplay, *errs.BError) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	list, err := d.projectAuthRepo.FindAuthList(c, orgCode)
	if err != nil {
		zap.L().Error("project AuthList projectAuthRepo.FindAuthList error", zap.Error(err))
		return nil, model.DBError
	}
	var pdList []*account.ProjectAuthDisplay
	for _, v := range list {
		display := v.ToDisplay()
		pdList = append(pdList, display)
	}
	return pdList, nil
}
