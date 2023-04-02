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
	group := r.Group("/project")
	//使用token认证中间件
	group.Use(midd.TokenVerify())
	group.POST("/index", h.index)
	group.POST("/project/selfList", h.myProjectList)
	group.POST("/project", h.myProjectList)
	group.POST("/project_template", h.projectTemplate)
}
