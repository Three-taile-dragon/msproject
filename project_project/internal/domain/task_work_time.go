package domain

import (
	"context"
	"go.uber.org/zap"
	"test.com/project_common/errs"
	"test.com/project_project/internal/dao/mysql"
	"test.com/project_project/internal/data"
	"test.com/project_project/internal/repo"
	"test.com/project_project/pkg/model"
	"time"
)

type TaskWorkTimeDomain struct {
	taskWorkTimeRepo repo.TaskWorkTimeRepo
	userRpcDomain    *UserRpcDomain
}

func NewTaskWorkTimeDomain() *TaskWorkTimeDomain {
	return &TaskWorkTimeDomain{
		taskWorkTimeRepo: mysql.NewTaskWorkTimeDao(),
		userRpcDomain:    NewUserRpcDomain(),
	}
}

func (d *TaskWorkTimeDomain) TaskWorkTimeList(taskCode int64) ([]*data.TaskWorkTimeDisplay, *errs.BError) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var list []*data.TaskWorkTime
	var err error
	list, err = d.taskWorkTimeRepo.FindWorkTimeList(c, taskCode)
	if err != nil {
		zap.L().Error("task TaskWorkTimeList taskWorkTimeRepo.FindWorkTimeList error", zap.Error(err))
		return nil, model.DBError
	}
	if len(list) == 0 {
		return []*data.TaskWorkTimeDisplay{}, nil
	}
	var displayList []*data.TaskWorkTimeDisplay
	var mIdList []int64

	for _, v := range list {
		mIdList = append(mIdList, v.MemberCode)
	}
	// 通过 domain 调用
	_, mMap, err := d.userRpcDomain.MemberList(mIdList)
	if err != nil {
		return nil, errs.ToBError(err)
	}
	for _, v := range list {
		display := v.ToDisplay()
		message := mMap[v.MemberCode]
		m := data.Member{}
		m.Name = message.Name
		m.Id = message.Id
		m.Avatar = message.Avatar
		m.Code = message.Code
		display.Member = m
		displayList = append(displayList, display)
	}
	return displayList, nil
}
