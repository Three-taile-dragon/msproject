package repo

import (
	"context"
	data "test.com/project_user/internal/data/organization"
)

type OrganizationRepo interface {
	FindOrganizationByMemId(ctx context.Context, memId int64) ([]data.Organization, error)
	SaveOrganization(ctx context.Context, org *data.Organization) error
}
