package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "test.com/project_api/api"
	"test.com/project_api/api/midd"
	"test.com/project_api/config"
	"test.com/project_api/router"
	srv "test.com/project_common"
)

func main() {
	r := gin.Default()

	r.Use(midd.RequestLog()) // 日志中间件

	// 静态目录映射
	r.StaticFS("/upload", http.Dir("upload"))

	router.InitRouter(r)
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, nil) //使用viper读取yaml配置文件
}
