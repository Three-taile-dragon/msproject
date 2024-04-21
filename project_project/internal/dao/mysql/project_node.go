package mysql

import (
	"context"
	"test.com/project_project/internal/data"
	"test.com/project_project/internal/database/gorms"
)

type ProjectNodeDao struct {
	conn *gorms.GormConn
}

func NewProjectNodeDao() *ProjectNodeDao {
	return &ProjectNodeDao{
		conn: gorms.New(),
	}
}

func (p *ProjectNodeDao) FindAll(ctx context.Context) (list []*data.ProjectNode, err error) {
	session := p.conn.Session(ctx)
	err = session.Model(&data.ProjectNode{}).Find(&list).Error
	return
}
