package project

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"test.com/project_api/api/midd"
	"test.com/project_api/api/rpc"
	"test.com/project_api/router"
)

type RouterProject struct {
}

func init() {
	log.Println("init project router")
	zap.L().Info("init project router")
	ru := &RouterProject{}
	router.Register(ru)
}

func (*RouterProject) Router(r *gin.Engine) {
	//初始化grpc的客户端连接
	rpc.InitRpcProjectClient()
	h := New()
	//路由组
	group := r.Group("/project/index")
	//使用token认证中间件
	group.Use(midd.TokenVerify())
	group.POST("", h.index)
	//路由组
	group1 := r.Group("/project/project")
	//使用token认证中间件
	group1.Use(midd.TokenVerify())
	group1.POST("/selfList", h.myProjectList)
	group1.POST("", h.myProjectList)
}
