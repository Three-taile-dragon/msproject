package domain

import (
	"context"
	"go.uber.org/zap"
	"test.com/project_common/errs"
	"test.com/project_project/internal/dao/mysql"
	"test.com/project_project/internal/data"
	"test.com/project_project/internal/repo"
	"test.com/project_project/pkg/model"
)

type ProjectNodeDomain struct {
	projectNodeRepo repo.ProjectNodeRepo
}

func NewProjectNodeDomain() *ProjectNodeDomain {
	return &ProjectNodeDomain{
		projectNodeRepo: mysql.NewProjectNodeDao(),
	}
}

func (d *ProjectNodeDomain) TreeList() ([]*data.ProjectNodeTree, *errs.BError) {
	// node 表都查出来 转换成 nodeList 结构
	nodes, err := d.projectNodeRepo.FindAll(context.Background())
	if err != nil {
		zap.L().Error("project ProjectNode TreeList projectNodeRepo.FindAll error", zap.Error(err))
		return nil, model.DBError
	}
	treeList := data.ToNodeTreeList(nodes)
	return treeList, nil
}

func (d *ProjectNodeDomain) AllNodeList() ([]*data.ProjectNode, *errs.BError) {
	list, err := d.projectNodeRepo.FindAll(context.Background())
	if err != nil {
		return nil, model.DBError
	}
	return list, nil
}
