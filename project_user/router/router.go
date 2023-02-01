package router

import (
	"github.com/gin-gonic/gin"
)

// Router 接口
type Router interface {
	Router(r *gin.Engine)
}

//type RegisterRouter struct {
//}

//func New() *RegisterRouter {
//	return &RegisterRouter{}
//}

//func (*RegisterRouter) Route(ro Router, r *gin.Engine) {
//	ro.Router(r)
//}

// routers 使用路由列表
// 后续添加路由，不需要再对InitRouter函数进行改动
var routers []Router

// InitRouter 初始化路由
func InitRouter(r *gin.Engine) {
	//reg := New()
	//路由注册
	//reg.Route(&user.RouteruUser{}, r)
	for _, reg := range routers {
		reg.Router(r)
	}

}

// Register 添加到路由列表中去
func Register(ro ...Router) {
	routers = append(routers, ro...)
}
