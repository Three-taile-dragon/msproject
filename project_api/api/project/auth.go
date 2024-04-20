package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"test.com/project_api/api/rpc"
	"test.com/project_api/pkg/model"
	common "test.com/project_common"
	"test.com/project_common/errs"
	account2 "test.com/project_grpc/account"
	"test.com/project_grpc/auth"
	"time"
)

type HandlerAuth struct{}

func NewAuth() *HandlerAuth {
	return &HandlerAuth{}
}

func (a HandlerAuth) auth(c *gin.Context) {
	// 解析参数
	result := &common.Result{}
	organizationCode := c.GetString("organizationCode")
	var page = &model.Page{}
	page.Bind(c)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	msg := &auth.AuthReqMessage{
		OrganizationCode: organizationCode,
		Page:             page.Page,
		PageSize:         page.PageSize,
	}

	response, err := rpc.AuthServiceClient.AuthList(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}

	var authList []*account2.ProjectAuth
	_ = copier.Copy(&authList, response.List)
	if authList == nil {
		authList = []*account2.ProjectAuth{}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"total": response.Total,
		"list":  authList,
		"page":  page.Page,
	}))
}
