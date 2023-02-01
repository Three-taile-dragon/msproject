package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()
	//r.Run(":8080")
	//写一个优雅的启停代码
	srv := &http.Server{
		Addr:    ":80",
		Handler: r,
	}
	//使用 go func 协程
	go func() {
		log.Printf("web server running is %s \n", srv.Addr)
		//srv.ListenAndServe 会阻塞线程
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()
	//接受关闭程序的信号
	quit := make(chan os.Signal)
	//SIGINT 用户发送INTR字符(Ctrl + C) 触发 kill -2
	//SIGTERM 结束程序(可以被捕获、阻塞或忽略)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting Down project web server...")
	//延时关闭
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Nanosecond)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("web server Shutdown, cause by : ", err)
	}
	select {
	case <-ctx.Done():
		log.Println("wait timeout...")
	}
	log.Println("web server stop success...")
}
