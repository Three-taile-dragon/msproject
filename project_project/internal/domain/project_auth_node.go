package domain

import (
	"context"
	"test.com/project_common/errs"
	"test.com/project_project/internal/dao/mysql"
	"test.com/project_project/internal/repo"
	"test.com/project_project/pkg/model"
	"time"
)

type ProjectAuthNodeDomain struct {
	projectAuthNodeRepo repo.ProjectAuthNodeRepo
}

func NewProjectAuthNodeDomain() *ProjectAuthNodeDomain {
	return &ProjectAuthNodeDomain{
		projectAuthNodeRepo: mysql.NewProjectAuthNodeDao(),
	}
}

func (an *ProjectAuthNodeDomain) AuthNodeList(authId int64) ([]string, *errs.BError) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	list, err := an.projectAuthNodeRepo.FindNodeStringList(c, authId)
	if err != nil {
		return nil, model.DBError
	}
	return list, nil
}
