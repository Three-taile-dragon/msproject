package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
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
	rsp, err := ProjectServiceClient.Index(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	c.JSON(http.StatusOK, result.Success(rsp.Menus))
}
