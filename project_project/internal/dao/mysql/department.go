package mysql

import (
	"context"
	"test.com/project_project/internal/data/department"
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

func (d *DepartmentDao) FindDepartmentById(ctx context.Context, id int64) (dt *department.Department, err error) {
	session := d.conn.Session(ctx)
	err = session.Where("id = ?", id).Find(&dt).Error
	return
}

func (d *DepartmentDao) ListDepartment(organizationCode int64, parentDepartmentCode int64, page int64, pageSize int64) (list []*department.Department, total int64, err error) {
	session := d.conn.Session(context.Background())
	session.Model(&department.Department{})
	session.Where("organization_code=?", organizationCode)
	if parentDepartmentCode > 0 {
		session.Where("pcode=?", parentDepartmentCode)
	}
	err = session.Count(&total).Error
	err = session.Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).Find(&list).Error
	return
}
