package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"test.com/project_api/api/rpc"
	"test.com/project_api/pkg/model/account"
	common "test.com/project_common"
	"test.com/project_common/errs"
	account2 "test.com/project_grpc/account"
	"time"
)

type HandlerAccount struct{}

func NewAccount() *HandlerAccount {
	return &HandlerAccount{}
}

func (a HandlerAccount) account(c *gin.Context) {
	// 1. 接受请求参数 一些参数的校验可以放在 api
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var req *account.AccountReq
	err2 := c.ShouldBind(&req)
	if err2 != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}

	// 2. 调用 project 模块 查询账户列表
	msg := &account2.AccountReqMessage{
		MemberId:         c.GetInt64("memberId"),
		OrganizationCode: c.GetString("organizationCode"),
		Page:             int64(req.Page),
		PageSize:         int64(req.PageSize),
		SearchType:       int32(req.SearchType),
		DepartmentCode:   req.DepartmentCode,
	}
	response, err := rpc.AccountServiceClient.Account(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	// 3. 返回数据
	var list []*account.MemberAccount
	_ = copier.Copy(&list, response.AccountList)
	if list == nil {
		list = []*account.MemberAccount{}
	}
	var authList []*account2.ProjectAuth
	_ = copier.Copy(&authList, response.AuthList)
	if authList == nil {
		authList = []*account2.ProjectAuth{}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"total":    response.Total,
		"page":     req.Page,
		"list":     list,
		"authList": authList,
	}))
}
