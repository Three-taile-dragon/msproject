package task_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"test.com/project_common/encrypts"
	"test.com/project_common/errs"
	"test.com/project_common/tms"
	"test.com/project_grpc/task"
	"test.com/project_project/internal/dao"
	"test.com/project_project/internal/dao/mysql"
	"test.com/project_project/internal/data"
	"test.com/project_project/internal/database/tran"
	"test.com/project_project/internal/repo"
	"test.com/project_project/pkg/model"
	"time"
)

// TaskService grpc 登陆服务 实现
type TaskService struct {
	task.UnimplementedTaskServiceServer
	cache                  repo.Cache
	transaction            tran.Transaction
	projectRepo            repo.ProjectRepo
	projectTemplateRepo    repo.ProjectTemplateRepo
	taskStagesTemplateRepo repo.TaskStagesTemplateRepo
	taskStagesRepo         repo.TaskStagesRepo
}

func New() *TaskService {
	return &TaskService{
		cache:                  dao.Rc,
		transaction:            dao.NewTransaction(),
		projectRepo:            mysql.NewProjectDao(),
		projectTemplateRepo:    mysql.NewProjectTemplateDao(),
		taskStagesTemplateRepo: mysql.NewTaskStagesTemplateDao(),
		taskStagesRepo:         mysql.NewTaskStagesDao(),
	}
}

func (t *TaskService) TaskStages(ctx context.Context, msg *task.TaskReqMessage) (*task.TaskStagesResponse, error) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	pmID := encrypts.DecryptNoErr(msg.ProjectCode)
	projectCode, err1 := t.projectRepo.FindProjectByCipId(c, pmID)
	if err1 != nil {
		zap.L().Error("task TaskStages projectRepo.FindProjectByCipId error", zap.Error(err1))
		return nil, errs.GrpcError(model.DBError)
	}

	page := msg.Page
	pageSIze := msg.PageSize

	stages, total, err := t.taskStagesRepo.FindStagesByProject(c, projectCode, page, pageSIze)
	if err != nil {
		zap.L().Error("task TaskStages taskStagesRepo.FindStagesByProject error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}

	var tsMessage []*task.TaskStagesMessage
	_ = copier.Copy(&tsMessage, &stages)
	if tsMessage == nil {
		return &task.TaskStagesResponse{List: tsMessage, Total: 0}, nil
	}
	stageMap := data.ToTaskStageMap(stages)
	for _, v := range tsMessage {
		taskStages := stageMap[int(v.Id)]
		v.Code = encrypts.EncryptInt64NoErr(int64(v.Id))
		v.CreateTime = tms.FormatByMill(taskStages.CreateTime)
		v.ProjectCode = encrypts.EncryptInt64NoErr(projectCode)
	}
	return &task.TaskStagesResponse{List: tsMessage, Total: total}, nil
}
