package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"net/http"
	"test.com/project_api/pkg/model/user"
	common "test.com/project_common"
	"test.com/project_common/errs"
	"test.com/project_grpc/user/login"
	"test.com/project_user/pkg/model"
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
	rsp, err := LoginServiceClient.GetCaptcha(c, &login.CaptchaRequest{Mobile: mobile})
	// 结果返回
	if err != nil {
		code, msg := errs.ParseGrpcError(err) //解析grpc错误信息
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(rsp.Code))
}

func (u *HandleUser) register(c *gin.Context) {
	//1.接收参数 参数模型
	result := &common.Result{}
	var req user.RegisterReq
	err := c.ShouldBind(&req) //绑定模型
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	//2.校验参数 判断参数是否合法
	if err := req.Verify(); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, err.Error()))
		return
	}
	//3.调用user grpc服务 获取响应
	//设置超时时间2秒
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	//使用 copier github.com/jinzhu/copier 来进行结构体赋值
	msg := &login.RegisterRequest{}
	err = copier.Copy(msg, req) //结构体赋值
	if err != nil {
		zap.L().Error("注册模块结构体赋值出错")
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "系统内部错误"))
		return
	}
	_, err = LoginServiceClient.Register(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err) //解析grpc错误信息
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	//4.返回结果
	c.JSON(http.StatusOK, result.Success(""))
}

func (u *HandleUser) login(c *gin.Context) {
	//1.接收参数
	result := &common.Result{}
	var req user.LoginReq
	err := c.ShouldBind(&req) //绑定模型
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	//2.校验参数 校验用户名是否存在
	if err := req.VerifyAccount(); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, err.Error()))
		return
	}
	//3.调用grpc服务 执行逻辑代码
	//设置超时时间2秒
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &login.LoginRequest{}
	err = copier.Copy(msg, req)
	if err != nil {
		zap.L().Error("登陆模块结构体赋值出错")
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "系统内部错误"))
		return
	}
	loginRsp, err := LoginServiceClient.Login(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err) //解析grpc错误信息
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	//4.返回结果
	//var rsp *user.LoginRsp
	//err = copier.Copy(&rsp, loginRsp)
	rsp := &user.LoginRsp{}
	err = copier.Copy(rsp, loginRsp)
	if err != nil {
		zap.L().Error("登陆模块orgs赋值错误", zap.Error(err))
		c.JSON(http.StatusOK, result.Fail(errs.ParseGrpcError(errs.GrpcError(model.SystemError))))
	}
	c.JSON(http.StatusOK, result.Success(rsp))
	//4.结果返回
}
