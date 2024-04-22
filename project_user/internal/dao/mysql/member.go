package mysql

import (
	"context"
	"errors"
	"gorm.io/gorm"
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

func (m MemberDao) GetMemberByAccountAndEmail(ctx context.Context, account string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("email=? or account=?", account, account).Count(&count).Error //数据库查询
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
func (m MemberDao) FindMember(ctx context.Context, account string, pwd string) (*member.Member, error) {
	var mem *member.Member
	err := m.conn.Session(ctx).Where("account=? and password=?", account, pwd).First(&mem).Error
	if err == gorm.ErrRecordNotFound {
		//未查询到对应的信息
		return nil, nil
	}
	return mem, err
}

func (m *MemberDao) FindMemberById(ctx context.Context, id int64) (*member.Member, error) {
	var mem *member.Member
	err := m.conn.Session(ctx).Where("id=?", id).First(&mem).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//未查询到对应的信息
		return nil, nil
	}
	return mem, err
}

func (m *MemberDao) FindMemberByIds(ctx context.Context, mIds []int64) (list []*member.Member, err error) {
	if len(mIds) <= 0 {
		return nil, nil
	}
	err = m.conn.Session(ctx).Model(&member.Member{}).Where("id in (?)", mIds).Find(&list).Error
	return
}
