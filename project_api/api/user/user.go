package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	common "test.com/project_common"
	"test.com/project_common/errs"
	loginservicev1 "test.com/project_user/pkg/service/login.service.v1"
	"time"
)

type HandleUser struct {
}

func New() *HandleUser {
	return &HandleUser{}
}

// 返回验证码
func (*HandleUser) getCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	// 获取传入的手机号
	mobile := ctx.PostForm("mobile")
	// 对grpc通信作允许两秒超时处理
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 通过grpc调用验证码生成函数
	rsp, err := LoginServiceClient.GetCaptcha(c, &loginservicev1.CaptchaRequest{Mobile: mobile})
	// 结果返回
	if err != nil {
		code, msg := errs.ParseGrpcError(err) //解析grpc错误信息
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(rsp.Code))
}
