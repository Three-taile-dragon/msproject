package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"test.com/project_user/router"
)

// 接口的实现类
// init 将user的路由 注册进入路由列表
func init() {
	log.Println("init user router")
	router.Register(&RouteruUser{})
}

type RouteruUser struct {
}

func (*RouteruUser) Router(r *gin.Engine) {
	h := New()
	r.POST("/project/login/getCaptcha", h.getCaptcha)
}
