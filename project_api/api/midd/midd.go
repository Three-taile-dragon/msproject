package midd

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"test.com/project_api/api/rpc"
	common "test.com/project_common"
	"test.com/project_common/errs"
	"test.com/project_grpc/user/login"
	"time"
)

func TokenVerify() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		//1.从Header中获取token
		result := &common.Result{}
		token := c.GetHeader("Authorization")
		//2.调用user服务进行token认证
		ctxo, canel := context.WithTimeout(context.Background(), 2*time.Second)
		defer canel()

		ip := GetIp(c)

		response, err := rpc.LoginServiceClient.TokenVerify(ctxo, &login.TokenRequest{Token: token, Ip: ip})
		//3.处理结果 认证通过，将信息放入gin上下文 失败就返回未登录
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(http.StatusOK, result.Fail(code, msg))
			c.Abort() //防止继续执行
			return
		}
		//成功
		c.Set("memberId", response.Member.Id)
		c.Set("memberName", response.Member.Name)
		c.Set("organizationCode", response.Member.OrganizationCode)
		c.Set("projectCode", response)
		c.Next()
	}
}

// GetIp 获取 ip 函数
func GetIp(c *gin.Context) string {
	ip := c.ClientIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	return ip
}
