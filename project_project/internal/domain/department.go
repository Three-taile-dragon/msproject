package domain

import (
	"context"
	"test.com/project_project/internal/dao/mysql"
	"test.com/project_project/internal/data/account"
	"test.com/project_project/internal/repo"
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

func (d *DepartmentDomain) FindDepartmentById(id int64) (*account.Department, error) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return d.departmentRepo.FindDepartmentById(c, id)
}
