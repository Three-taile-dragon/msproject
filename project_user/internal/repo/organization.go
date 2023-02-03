package repo

import (
	"context"
	data "test.com/project_user/internal/data/organization"
	"test.com/project_user/internal/database"
)

type OrganizationRepo interface {
	FindOrganizationByMemId(ctx context.Context, memId int64) ([]data.Organization, error)
	SaveOrganization(conn database.DbConn, ctx context.Context, org *data.Organization) error
}
