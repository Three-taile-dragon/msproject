package repo

import (
	"context"
	"test.com/project_project/internal/data"
	"test.com/project_project/internal/database"
)

type TaskStagesTemplateRepo interface {
	FindInProTemIds(ctx context.Context, id []int) ([]data.MsTaskStagesTemplate, error)
	FindByProjectTemplateId(ctx context.Context, projectTemplateCode int) (list []*data.MsTaskStagesTemplate, err error)
}

type TaskStagesRepo interface {
	SaveTaskStages(ctx context.Context, conn database.DbConn, ts *data.TaskStages) error
	FindStagesByProject(ctx context.Context, projectCode int64, page int64, pageSize int64) (list []*data.TaskStages, total int64, err error)
}

type TaskRepo interface {
	FindTaskByStageCode(ctx context.Context, stageCode int) (list []*data.Task, err error)
}

type TaskMemberRepo interface {
	FindTaskMemberByTaskId(ctx context.Context, taskCode int64, memberId int64) (task *data.TaskMember, err error)
}
