package main

import (
	"github.com/gin-gonic/gin"
	"log"
	srv "test.com/project_common"
	"test.com/project_common/logs"
	_ "test.com/project_user/api"
	"test.com/project_user/router"
)

func main() {
	r := gin.Default()
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: "F:\\Go_item\\ms_project\\logs\\debug\\project-debug.log",
		InfoFileName:  "F:\\Go_item\\ms_project\\logs\\info\\project-info.log",
		WarnFileName:  "F:\\Go_item\\ms_project\\logs\\error\\project-error.log",
		MaxSize:       500,
		MaxAge:        28,
		MaxBackups:    3,
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
	r.Use(logs.GinLogger(), logs.GinRecovery(true)) //接收gin框架默认日志
	//路由
	router.InitRouter(r)
	//r.Run(":8080")
	srv.Run(r, "project_user", ":3456")
}
