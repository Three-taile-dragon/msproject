package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"test.com/project_api/api/midd"
	"test.com/project_api/api/rpc"
	"test.com/project_api/router"
)

type RouterUser struct {
}

func init() {
	log.Println("init user router")
	zap.L().Info("init user router")
	ru := &RouterUser{}
	router.Register(ru)
}

func (*RouterUser) Router(r *gin.Engine) {
	//初始化grpc的客户端连接
	rpc.InitRpcUserClient()
	h := New()
	r.POST("/project/login/getCaptcha", h.getCaptcha)
	r.POST("/project/login/register", h.register)
	r.POST("/project/login", h.login)
	org := r.Group("/project/organization") //创建组
	org.Use(midd.TokenVerify())             //Token认证
	org.POST("/_getOrgList", h.myOrgList)
}
