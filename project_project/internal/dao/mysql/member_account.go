package mysql

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"test.com/project_project/internal/data/account"
	"test.com/project_project/internal/database/gorms"
)

type MemberAccountDao struct {
	conn *gorms.GormConn
}

func NewMemberAccountDao() *MemberAccountDao {
	return &MemberAccountDao{
		conn: gorms.New(),
	}
}

func (m *MemberAccountDao) FindList(ctx context.Context, condition string, organizationCode int64, departmentCode int64, page int64, pageSize int64) (list []*account.MemberAccount, total int64, err error) {
	session := m.conn.Session(ctx)
	offset := (page - 1) * pageSize
	err = session.Model(&account.MemberAccount{}).
		Where("organization_code=?", organizationCode).
		Where(condition).Limit(int(pageSize)).Offset(int(offset)).Find(&list).Error
	err = session.Model(&account.MemberAccount{}).
		Where("organization_code=?", organizationCode).
		Where(condition).Count(&total).Error
	return
}

func (m *MemberAccountDao) FindByMemberId(ctx context.Context, memberId int64) (ma *account.MemberAccount, err error) {
	session := m.conn.Session(ctx)
	err = session.Where("member_code = ?", memberId).Take(&ma).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}
