package domain

import (
	"context"
	"test.com/project_common/errs"
	"test.com/project_project/internal/dao/mysql"
	"test.com/project_project/internal/data/department"
	"test.com/project_project/internal/repo"
	"test.com/project_project/pkg/model"
	"time"
)

type DepartmentDomain struct {
	departmentRepo repo.DepartmentRepo
}

func NewDepartmentDomain() *DepartmentDomain {
	return &DepartmentDomain{
		departmentRepo: mysql.NewDepartmentDao(),
	}
}

func (d *DepartmentDomain) FindDepartmentById(id int64) (*department.Department, error) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return d.departmentRepo.FindDepartmentById(c, id)
}

func (d *DepartmentDomain) List(organizationCode int64, departmentCode int64, page int64, pageSize int64) ([]*department.DepartmentDisplay, int64, *errs.BError) {
	list, total, err := d.departmentRepo.ListDepartment(organizationCode, departmentCode, page, pageSize)
	if err != nil {
		return nil, 0, model.DBError
	}
	var dList []*department.DepartmentDisplay
	for _, v := range list {
		dList = append(dList, v.ToDisplay())
	}
	return dList, total, nil
}
