package task_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"test.com/project_common/encrypts"
	"test.com/project_common/errs"
	"test.com/project_common/tms"
	"test.com/project_grpc/task"
	"test.com/project_grpc/user/login"
	"test.com/project_project/internal/dao"
	"test.com/project_project/internal/dao/mysql"
	"test.com/project_project/internal/data"
	pro "test.com/project_project/internal/data/project"
	"test.com/project_project/internal/database/tran"
	"test.com/project_project/internal/repo"
	"test.com/project_project/internal/rpc"
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

func (t *TaskService) MemberProjectList(ctx context.Context, msg *task.TaskReqMessage) (*task.MemberProjectResponse, error) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 查询 用户id 列表
	pmID := encrypts.DecryptNoErr(msg.ProjectCode)
	projectCode, err1 := t.projectRepo.FindProjectByCipId(c, pmID)
	if err1 != nil {
		zap.L().Error("task TaskStages projectRepo.FindProjectByCipId error", zap.Error(err1))
		return nil, errs.GrpcError(model.DBError)
	}
	projectMembers, total, err := t.projectRepo.FindProjectByPid(c, projectCode)
	if err != nil {
		zap.L().Error("task MemberProjectList projectRepo.FindProjectByPid error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	// 使用 用户id 列表 请求用户信息
	if projectMembers == nil || len(projectMembers) <= 0 {
		return &task.MemberProjectResponse{List: nil, Total: 0}, nil
	}
	var mIds []int64
	pmMap := make(map[int64]*pro.ProjectMember)
	for _, v := range projectMembers {
		mIds = append(mIds, v.MemberCode)
		pmMap[v.MemberCode] = v
	}

	//请求用户信息
	userMsg := &login.UserMessage{
		MIds: mIds,
	}
	memberMessageList, err := rpc.LoginServiceClient.FindMemInfoByIds(ctx, userMsg)
	if err != nil {
		zap.L().Error("task MemberProjectList LoginServiceClient.FindMemInfoByIds error", zap.Error(err))
		return nil, err
	}

	var list []*task.MemberProjectMessage
	for _, v := range memberMessageList.List {
		owner := pmMap[v.Id].IsOwner
		mpm := &task.MemberProjectMessage{
			MemberCode: v.Id,
			Name:       v.Name,
			Avatar:     v.Avatar,
			Email:      v.Email,
			Code:       v.Code,
			IsOwner:    model.NoOwner,
		}
		if v.Id == owner {
			mpm.IsOwner = model.Owner
		}
		list = append(list, mpm)
	}

	return &task.MemberProjectResponse{List: list, Total: total}, nil

}
