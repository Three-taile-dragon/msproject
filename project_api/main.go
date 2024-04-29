package main

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"log"
	"net/http"
	_ "test.com/project_api/api"
	"test.com/project_api/api/midd"
	"test.com/project_api/config"
	"test.com/project_api/router"
	"test.com/project_api/tracing"
	srv "test.com/project_common"
)

func main() {
	r := gin.Default()
	// 加载 链路追踪 jaeger
	tp, tpErr := tracing.JaegerTraceProvider()
	if tpErr != nil {
		log.Fatal(tpErr)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	r.Use(midd.RequestLog()) // 日志中间件

	r.Use(otelgin.Middleware("msProject-api")) // 使用中间件形式 加载 插件

	// 静态目录映射
	r.StaticFS("/upload", http.Dir("upload"))

	router.InitRouter(r)

	// 开启 pprof 默认访问路径 /debug/pprof
	//pprof.Register(r, "/dev/pprof")

	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, nil) //使用viper读取yaml配置文件
}
