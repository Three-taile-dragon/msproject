package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"test.com/project_api/api/rpc"
	"test.com/project_api/pkg/model/menu"
	common "test.com/project_common"
	"test.com/project_common/errs"
	menu2 "test.com/project_grpc/menu"
	"time"
)

type HandlerMenu struct{}

func NewMenu() *HandlerMenu {
	return &HandlerMenu{}
}

func (m *HandlerMenu) menuList(c *gin.Context) {
	// 直接返回结果即可
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := rpc.MenuServiceClient.MenuList(ctx, &menu2.MenuReqMessage{})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	var list []*menu.Menu
	_ = copier.Copy(&list, res.List)
	if list == nil {
		list = []*menu.Menu{}
	}
	c.JSON(http.StatusOK, result.Success(list))
}
