package repo

import (
	"context"
	"test.com/project_project/internal/data/account"
)

type AccountRepo interface {
	FindList(ctx context.Context, condition string, organizationCode int64, departmentCode int64, page int64, pageSize int64) (list []*account.MemberAccount, total int64, err error)
	FindByMemberId(ctx context.Context, memberId int64) (ma *account.MemberAccount, err error)
}
