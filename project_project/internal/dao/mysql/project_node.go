package mysql

import (
	"context"
	"test.com/project_project/internal/data"
	"test.com/project_project/internal/database"
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

func (p *ProjectNodeDao) DeleteByAuthId(ctx context.Context, conn database.DbConn, authId int64) error {
	p.conn = conn.(*gorms.GormConn)
	tx := p.conn.Tx(ctx)
	err := tx.Where("auth=?", authId).Delete(&data.ProjectAuthNode{}).Error
	return err
}

func (p *ProjectNodeDao) Save(ctx context.Context, conn database.DbConn, authId int64, nodes []string) error {
	p.conn = conn.(*gorms.GormConn)
	tx := p.conn.Tx(ctx)
	var list []*data.ProjectAuthNode
	for _, v := range nodes {
		pn := &data.ProjectAuthNode{
			Auth: authId,
			Node: v,
		}
		list = append(list, pn)
	}
	err := tx.Create(list).Error
	return err
}
