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
	FindProjectByPIdAndMemId(ctx context.Context, projectCode int64, memberId int64) (*project.ProjectAndMember, error)
	FindCollectByPIdAndMemId(ctx context.Context, projectCode int64, memberId int64) (bool, error)
	FindProjectByCipId(ctx context.Context, cipherIdCode int64) (int64, error)
	DeleteProject(ctx context.Context, id int64) error
	RecoveryProject(ctx context.Context, id int64) error
	CollectProject(ctx context.Context, pc *project.ProjectCollection) error
	CancelCollectProject(ctx context.Context, projectCode int64, memberId int64) error
}

// ProjectTemplateRepo 查询项目模板
type ProjectTemplateRepo interface {
	FindProjectTemplateSystem(ctx context.Context, page int64, size int64) ([]project.ProjectTemplate, int64, error)
	FindProjectTemplateCustom(ctx context.Context, memId int64, organizationCode int64, page int64, size int64) ([]project.ProjectTemplate, int64, error)
	FindProjectTemplateAll(ctx context.Context, organizationCode int64, page int64, size int64) ([]project.ProjectTemplate, int64, error)
}
