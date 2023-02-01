package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	common "test.com/project_common"
	"test.com/project_user/pkg/dao"
	"test.com/project_user/pkg/model"
	"test.com/project_user/pkg/repo"
	"test.com/project_user/util"
	"time"
)

type HandlerUser struct {
	cache repo.Cache
}

func New() *HandlerUser {
	return &HandlerUser{
		cache: dao.Rc, //使用redis
	}
}

func (h *HandlerUser) getCaptcha(ctx *gin.Context) {
	rsp := &common.Result{}
	//1.获取参数
	mobile := ctx.PostForm("mobile")
	//2.校验参数
	if !common.VerifyMobile(mobile) {
		// 2001 这里为魔法数字	指无法第一眼得出该数字所代表的含义
		//应当避免魔法数字
		ctx.JSON(http.StatusOK, rsp.Fail(model.NoLegalMobile, "手机号不合法"))
		return
	}
	//3.生成验证码(随机四位1000-9999或者六位100000-999999)
	code := util.CreateCaptcha(6) //生成随机六位数字验证码
	//4.调用短信平台(第三方 放入go func 协程 接口可以快速响应
	go func() {
		time.Sleep(2 * time.Second)
		zap.L().Info("短信平台调用成功，发送短信")
		//logs.LG.Debug("短信平台调用成功，发送短信 debug")
		//zap.L().Debug("短信平台调用成功，发送短信 debug")
		//zap.L().Error("短信平台调用成功，发送短信 error")
		//redis存储	假设后续缓存可能存在mysql当中,也可以存在mongo当中,也可能存在memcache当中
		//使用接口 达到低耦合高内聚
		//5.存储验证码 redis 当中,过期时间15分钟
		//redis.Set"REGISTER_"+mobile, code)
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second) //编写上下文 最多允许两秒超时
		defer cancel()
		err := h.cache.Put(c, "REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			zap.L().Error("验证码存入redis出错,cause by : " + err.Error() + "\n")

		}
		zap.L().Debug("将手机号和验证码存入redis成功：REGISTER_" + mobile + " : " + code + "\n")
	}()
	//注意code一般不发送
	//这里是做了简化处理 由于短信平台目前对于个人不好使用
	ctx.JSON(http.StatusOK, rsp.Success(code))
}
