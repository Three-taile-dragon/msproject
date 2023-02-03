package repo

import (
	"context"
	"test.com/project_user/internal/data/member"
)

type MemberRepo interface {
	SaveMember(ctx context.Context, mem *member.Member) error
	GetMemberByEmail(ctx context.Context, email string) (bool, error)
	GetMemberByAccount(ctx context.Context, account string) (bool, error)
	GetMemberByName(ctx context.Context, name string) (bool, error)
	GetMemberByMobile(ctx context.Context, mobile string) (bool, error)
}
