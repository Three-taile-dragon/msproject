package repo

import (
	"context"
	"test.com/project_project/internal/data/department"
)

type DepartmentRepo interface {
	FindDepartmentById(ctx context.Context, id int64) (dt *department.Department, err error)
	ListDepartment(organizationCode int64, parentDepartmentCode int64, page int64, pageSize int64) (list []*department.Department, total int64, err error)
	FindDepartment(ctx context.Context, organizationCode int64, parentDepartmentCode int64, name string) (*department.Department, error)
	Save(dpm *department.Department) error
}
