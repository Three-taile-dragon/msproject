package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	common "test.com/project_common"
	"test.com/project_user/pkg/model"
	"test.com/project_user/util"
	"time"
)

type HandlerUser struct {
}

func (*HandlerUser) getCaptcha(ctx *gin.Context) {
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
		log.Println("短信平台调用成功，发送短信")
		//5.存储验证码 redis 当中,过期时间15分钟
		log.Printf("将手机号和验证码存入redis成功：REGISTER_%s : %s", mobile, code)
	}()
	//注意code一般不发送
	//这里是做了简化处理 由于短信平台目前对于个人不好使用
	ctx.JSON(http.StatusOK, rsp.Success(code))
}
