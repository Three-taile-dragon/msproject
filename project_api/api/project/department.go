package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"test.com/project_api/api/rpc"
	"test.com/project_api/pkg/model/department"
	common "test.com/project_common"
	"test.com/project_common/errs"
	department2 "test.com/project_grpc/department"
	"time"
)

type HandlerDepartment struct{}

func NewDepartment() *HandlerDepartment {
	return &HandlerDepartment{}
}

func (d *HandlerDepartment) department(c *gin.Context) {
	// 解析参数
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var req *department.DepartmentReq
	err2 := c.ShouldBind(&req)
	if err2 != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}

	// 调用 rpc 服务
	msg := &department2.DepartmentReqMessage{
		Page:                 req.Page,
		PageSize:             req.PageSize,
		ParentDepartmentCode: req.Pcode,
		OrganizationCode:     c.GetString("organizationCode"),
	}
	listDepartmentMessage, err := rpc.DepartmentServiceClient.List(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	// 返回数据
	var list []*department.Department
	_ = copier.Copy(&list, listDepartmentMessage.List)
	if list == nil {
		list = []*department.Department{}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"total": listDepartmentMessage.Total,
		"page":  req.Page,
		"list":  list,
	}))

}

func (d *HandlerDepartment) save(c *gin.Context) {
	result := &common.Result{}
	var req *department.DepartmentReq
	err2 := c.ShouldBind(&req)
	if err2 != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &department2.DepartmentReqMessage{
		Name:                 req.Name,
		DepartmentCode:       req.DepartmentCode,
		ParentDepartmentCode: req.ParentDepartmentCode,
		OrganizationCode:     c.GetString("organizationCode"),
	}
	departmentMessage, err := rpc.DepartmentServiceClient.Save(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	var res = &department.Department{}
	_ = copier.Copy(res, departmentMessage)
	c.JSON(http.StatusOK, result.Success(res))
}
