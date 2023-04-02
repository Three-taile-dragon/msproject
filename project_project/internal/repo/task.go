package repo

import (
	"context"
	"test.com/project_project/internal/data/task"
)

type TaskStagesTemplateRepo interface {
	FindInProTemIds(ctx context.Context, id []int) ([]task.MsTaskStagesTemplate, error)
}
