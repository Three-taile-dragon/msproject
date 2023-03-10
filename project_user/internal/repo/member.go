package repo

import (
	"context"
	"test.com/project_user/internal/data/member"
	"test.com/project_user/internal/database"
)

type MemberRepo interface {
	SaveMember(conn database.DbConn, ctx context.Context, mem *member.Member) error
	GetMemberByEmail(ctx context.Context, email string) (bool, error)
	GetMemberByAccount(ctx context.Context, account string) (bool, error)
	GetMemberByAccountAndEmail(ctx context.Context, account string) (bool, error)
	GetMemberByName(ctx context.Context, name string) (bool, error)
	GetMemberByMobile(ctx context.Context, mobile string) (bool, error)
	FindMember(ctx context.Context, account string, pwd string) (mem *member.Member, err error)
	FindMemberById(ctx context.Context, id int64) (mem *member.Member, err error)
}
