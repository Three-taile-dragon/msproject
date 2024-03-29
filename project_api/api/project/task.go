package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"net/http"
	"test.com/project_api/api/rpc"
	"test.com/project_api/pkg/model"
	"test.com/project_api/pkg/model/tasks"
	common "test.com/project_common"
	"test.com/project_common/errs"
	"test.com/project_grpc/task"
	"time"
)

type HandlerTask struct {
}

func NewTask() *HandlerTask {
	return &HandlerTask{}
}

func (t HandlerTask) taskStages(c *gin.Context) {
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	//获取参数	校验参数合法性
	projectCode := c.PostForm("projectCode")
	page := &model.Page{}
	page.Bind(c)
	//调用grpc
	msg := &task.TaskReqMessage{
		MemberId:    c.GetInt64("memberId"),
		ProjectCode: projectCode,
		Page:        page.Page,
		PageSize:    page.PageSize,
	}
	stages, err := rpc.TaskServiceClient.TaskStages(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	//处理响应
	var list []*tasks.TaskStagesResp
	err = copier.Copy(&list, stages.List)
	if list == nil {
		list = []*tasks.TaskStagesResp{}
	}
	for _, v := range list {
		v.TasksLoading = true  //任务加载状态
		v.FixedCreator = false //添加任务按钮定位
		v.ShowTaskCard = false //是否显示创建卡片
		v.Tasks = []int{}
		v.DoneTasks = []int{}
		v.UnDoneTasks = []int{}
	}
	if err != nil {
		zap.L().Error("task taskStages TaskStagesResponse Copy error", zap.Error(err))
		c.JSON(http.StatusOK, result.Fail(errs.ParseGrpcError(err)))
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  list,
		"total": stages.Total,
		"page":  page.Page,
	}))

}
