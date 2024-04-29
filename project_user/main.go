package main

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"log"
	srv "test.com/project_common"
	"test.com/project_user/config"
	"test.com/project_user/router"
	"test.com/project_user/tracing"
)

func main() {
	r := gin.Default()
	//从配置中读取日志配置，初始化日志
	// 加载 链路追踪 jaeger
	tp, tpErr := tracing.JaegerTraceProvider()
	if tpErr != nil {
		log.Fatal(tpErr)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	//r.Use(logs.GinLogger(), logs.GinRecovery(true)) //接收gin框架默认日志
	r.Use(otelgin.Middleware("msProject-user")) // 使用中间件形式 加载 插件
	//路由
	router.InitRouter(r)
	//grpc注册
	gc := router.RegisterGrpc()
	//grpc 服务注册到etcd
	router.RegisterEtcdServer()
	stop := func() {
		gc.Stop()
	}
	//r.Run(":8080")
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, stop) //使用viper读取yaml配置文件
}
