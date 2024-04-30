package domain

import (
	"context"
	"test.com/project_common/errs"
	"test.com/project_common/kk"
	"test.com/project_project/config"
	"test.com/project_project/internal/dao/mysql"
	"test.com/project_project/internal/repo"
	"test.com/project_project/pkg/model"
)

type TaskDomain struct {
	taskRepo repo.TaskRepo
}

func NewTaskDomain() *TaskDomain {
	return &TaskDomain{
		taskRepo: mysql.NewTaskDao(),
	}
}

func (d *TaskDomain) FindProjectIdByTaskId(taskId int64) (int64, bool, *errs.BError) {
	// 记录日志
	config.SendLog(kk.Info("Find", "TaskDomain.FindProjectIdByTaskId", kk.FieldMap{
		"taskId": taskId,
	}))

	task, err := d.taskRepo.FindTaskById(context.Background(), taskId)
	if err != nil {
		// 记录错误日志
		config.SendLog(kk.Error(err, "TaskDomain.FindProjectIdByTaskId.taskRepo.FindTaskById", kk.FieldMap{
			"taskId": taskId,
		}))
		return 0, false, model.DBError
	}
	if task == nil {
		return 0, false, nil
	}
	return task.ProjectCode, true, nil
}
