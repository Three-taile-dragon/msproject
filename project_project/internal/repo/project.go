package repo

import (
	"context"
	"test.com/project_project/internal/data/project"
	"test.com/project_project/internal/database"
)

// ProjectRepo 查询项目
type ProjectRepo interface {
	FindProjectByMemId(ctx context.Context, memId int64, condition string, page int64, size int64) ([]*project.ProjectAndMember, int64, error)
	FindCollectProjectByMemId(ctx context.Context, memberId int64, page int64, size int64) ([]*project.ProjectAndMember, int64, error)
	SaveProject(conn database.DbConn, ctx context.Context, pr *project.Project) error
	SaveProjectMember(conn database.DbConn, ctx context.Context, pm *project.ProjectMember) error
}

// ProjectTemplateRepo 查询项目模板
type ProjectTemplateRepo interface {
	FindProjectTemplateSystem(ctx context.Context, page int64, size int64) ([]project.ProjectTemplate, int64, error)
	FindProjectTemplateCustom(ctx context.Context, memId int64, organizationCode int64, page int64, size int64) ([]project.ProjectTemplate, int64, error)
	FindProjectTemplateAll(ctx context.Context, organizationCode int64, page int64, size int64) ([]project.ProjectTemplate, int64, error)
}
