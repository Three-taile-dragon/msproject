package mysql

import (
	"context"
	"test.com/project_project/internal/data"
	"test.com/project_project/internal/database/gorms"
)

type ProjectAuthNodeDao struct {
	conn *gorms.GormConn
}

func (p *ProjectAuthNodeDao) FindNodeStringList(ctx context.Context, authId int64) (list []string, err error) {
	session := p.conn.Session(ctx)
	err = session.Model(&data.ProjectAuthNode{}).Where("auth=?", authId).Select("node").Find(&list).Error
	return
}

func NewProjectAuthNodeDao() *ProjectAuthNodeDao {
	return &ProjectAuthNodeDao{
		conn: gorms.New(),
	}
}
