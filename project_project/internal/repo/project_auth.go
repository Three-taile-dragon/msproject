package repo

import (
	"context"
	"test.com/project_project/internal/data/account"
)

type ProjectAuthRepo interface {
	FindAuthList(ctx context.Context, orgCode int64) (list []*account.ProjectAuth, err error)
	FindAuthListPage(ctx context.Context, organizationCode int64, page int64, pageSize int64) (list []*account.ProjectAuth, total int64, err error)
}
