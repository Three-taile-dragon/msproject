package repo

import (
	"context"
	"test.com/project_project/internal/data/account"
)

type DepartmentRepo interface {
	FindDepartmentById(ctx context.Context, id int64) (dt *account.Department, err error)
}
