package mysql

import (
	"context"
	"test.com/project_user/internal/data/member"
	"test.com/project_user/internal/database"
	"test.com/project_user/internal/database/gorms"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func NewMemberDao() *MemberDao {
	return &MemberDao{
		conn: gorms.New(),
	}
}

func (m *MemberDao) SaveMember(conn database.DbConn, ctx context.Context, mem *member.Member) error {
	m.conn = conn.(*gorms.GormConn) //使用事务操作
	return m.conn.Tx(ctx).Create(mem).Error
}

func (m MemberDao) GetMemberByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("email=?", email).Count(&count).Error //数据库查询
	return count > 0, err
}

func (m MemberDao) GetMemberByAccount(ctx context.Context, account string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("account=?", account).Count(&count).Error //数据库查询
	return count > 0, err
}

func (m MemberDao) GetMemberByName(ctx context.Context, name string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("name=?", name).Count(&count).Error //数据库查询
	return count > 0, err
}

func (m MemberDao) GetMemberByMobile(ctx context.Context, mobile string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("mobile=?", mobile).Count(&count).Error //数据库查询
	return count > 0, err
}
