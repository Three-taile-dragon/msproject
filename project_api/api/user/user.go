package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	common "test.com/project_common"
	loginservicev1 "test.com/project_user/pkg/service/login.service.v1"
	"time"
)

type HandleUser struct {
}

func New() *HandleUser {
	return &HandleUser{}
}

func (*HandleUser) getCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	mobile := ctx.PostForm("mobile")
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	rsp, err := LoginServiceClient.GetCaptcha(c, &loginservicev1.CaptchaRequest{Mobile: mobile})
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(2001, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(rsp.Code))
}
