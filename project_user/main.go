package main

import (
	"github.com/gin-gonic/gin"
	srv "test.com/project_common"
	"test.com/project_common/logs"
	_ "test.com/project_user/api"
	"test.com/project_user/config"
	"test.com/project_user/router"
)

func main() {
	r := gin.Default()
	//从配置中读取日志配置，初始化日志

	r.Use(logs.GinLogger(), logs.GinRecovery(true)) //接收gin框架默认日志
	//路由
	router.InitRouter(r)
	//r.Run(":8080")
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr) //使用viper读取yaml配置文件
}
