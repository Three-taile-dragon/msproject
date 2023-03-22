package mysql

import (
	"context"
	data "test.com/project_user/internal/data/organization"
	"test.com/project_user/internal/database"
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

func (o *OrganizationDao) FindOrganizationByMemId(ctx context.Context, memId int64) ([]*data.Organization, error) {
	var orgs []*data.Organization
	err := o.conn.Session(ctx).Where("member_id=?", memId).Find(&orgs).Error
	return orgs, err
}

func (o *OrganizationDao) SaveOrganization(conn database.DbConn, ctx context.Context, org *data.Organization) error {
	o.conn = conn.(*gorms.GormConn) //使用事务操作
	return o.conn.Tx(ctx).Create(org).Error
}
