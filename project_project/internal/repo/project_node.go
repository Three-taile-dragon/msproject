package repo

import (
	"context"
	"test.com/project_project/internal/data"
	"test.com/project_project/internal/database"
)

type ProjectNodeRepo interface {
	FindAll(ctx context.Context) (list []*data.ProjectNode, err error)
	DeleteByAuthId(ctx context.Context, conn database.DbConn, authId int64) error
	Save(ctx context.Context, conn database.DbConn, authId int64, nodes []string) error
}
