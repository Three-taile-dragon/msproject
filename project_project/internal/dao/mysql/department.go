package mysql

import (
	"context"
	"test.com/project_project/internal/data/account"
	"test.com/project_project/internal/database/gorms"
)

type DepartmentDao struct {
	conn *gorms.GormConn
}

func NewDepartmentDao() *DepartmentDao {
	return &DepartmentDao{
		conn: gorms.New(),
	}
}

func (d *DepartmentDao) FindDepartmentById(ctx context.Context, id int64) (dt *account.Department, err error) {
	session := d.conn.Session(ctx)
	err = session.Where("id = ?", id).Find(&dt).Error
	return
}
