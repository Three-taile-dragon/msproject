package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"net/http"
	"test.com/project_api/api/rpc"
	"test.com/project_api/pkg/model"
	pro "test.com/project_api/pkg/model/project"
	common "test.com/project_common"
	"test.com/project_common/errs"
	"test.com/project_grpc/project"
	"time"
)

type HandleProject struct {
}

func New() *HandleProject {
	return &HandleProject{}
}
func (p *HandleProject) index(c *gin.Context) {
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &project.IndexRequest{}
	rsp, err := rpc.ProjectServiceClient.Index(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	c.JSON(http.StatusOK, result.Success(rsp.Menus))
}

func (p *HandleProject) myProjectList(c *gin.Context) {
	result := &common.Result{}
	//1. 获取参数
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberIdStr, _ := c.Get("memberId")
	memberId := memberIdStr.(int64) //转换
	//分页
	page := &model.Page{}
	page.Bind(c)
	msg := &project.ProjectRpcMessage{MemberId: memberId, Page: page.Page, PageSize: page.PageSize}
	myProjectResponse, err := rpc.ProjectServiceClient.FindProjectByMemId(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	//设定默认值
	if myProjectResponse.Pm == nil {
		myProjectResponse.Pm = []*project.ProjectMessage{}
	}
	var pms []*pro.ProjectAndMember
	err = copier.Copy(&pms, myProjectResponse)
	if err != nil {
		zap.L().Error("项目列表模块返回数据复制出错", zap.Error(err))
		c.JSON(http.StatusOK, result.Fail(502, "系统内部错误"))
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  pms,
		"total": myProjectResponse.Total,
	}))
}
