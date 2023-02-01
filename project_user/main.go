package main

import (
	"github.com/gin-gonic/gin"
	srv "test.com/project_common"
	_ "test.com/project_user/api"
	"test.com/project_user/router"
)

func main() {
	r := gin.Default()
	//路由
	router.InitRouter(r)
	//r.Run(":8080")
	srv.Run(r, "project_user", ":80")
}
