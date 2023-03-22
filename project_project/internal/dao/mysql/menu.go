package mysql

import (
	"context"
	"test.com/project_project/internal/data/menu"
	"test.com/project_project/internal/database/gorms"
)

type MenuDao struct {
	conn *gorms.GormConn
}

func NewMenuDao() *MenuDao {
	return &MenuDao{
		conn: gorms.New(),
	}
}

// FindMenus 实现菜单查询
func (m *MenuDao) FindMenus(ctx context.Context) (pms []*menu.ProjectMenu, err error) {
	session := m.conn.Session(ctx)
	err = session.Order("pid,sort asc, id asc").Find(&pms).Error
	return
}
