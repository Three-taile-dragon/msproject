package mysql

import (
	"context"
	"errors"
	"gorm.io/gorm"
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
	session = session.Model(&department.Department{})
	session = session.Where("organization_code=?", organizationCode)
	if parentDepartmentCode > 0 {
		session = session.Where("pcode=?", parentDepartmentCode)
	}
	err = session.Count(&total).Error
	err = session.Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).Find(&list).Error
	return
}

func (d *DepartmentDao) FindDepartment(ctx context.Context, organizationCode int64, parentDepartmentCode int64, name string) (*department.Department, error) {
	session := d.conn.Session(ctx)
	session = session.Model(&department.Department{}).Where("organization_code=? AND name=?", organizationCode, name)
	if parentDepartmentCode > 0 {
		session = session.Where("pcode=?", parentDepartmentCode)
	}
	var dp *department.Department
	err := session.Limit(1).Take(&dp).Error
	// 注意 dp 不能为 nil
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return dp, err
}

func (d *DepartmentDao) Save(dpm *department.Department) error {
	err := d.conn.Session(context.Background()).Save(&dpm).Error
	return err
}
