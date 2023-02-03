package mysql

import (
	"context"
	data "test.com/project_user/internal/data/organization"
	"test.com/project_user/internal/database/gorms"
)

type OrganizationDao struct {
	conn *gorms.GormConn
}

func NewOrganizationDao() *OrganizationDao {
	return &OrganizationDao{
		conn: gorms.New(),
	}
}

func (o *OrganizationDao) FindOrganizationByMemId(ctx context.Context, memId int64) ([]data.Organization, error) {
	var orgs []data.Organization
	err := o.conn.Session(ctx).Where("member_id=?", memId).Find(&orgs).Error
	return orgs, err
}

func (o *OrganizationDao) SaveOrganization(ctx context.Context, org *data.Organization) error {
	err := o.conn.Session(ctx).Create(org).Error
	return err
}
