package user

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	login_service_v1 "test.com/project_user/pkg/service/login.service.v1"
)

var LoginServiceClient login_service_v1.LoginServiceClient

// InitRpcUserClient 初始化grpc客户段连接
func InitRpcUserClient() {
	conn, err := grpc.Dial("127.0.0.1:8881", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	LoginServiceClient = login_service_v1.NewLoginServiceClient(conn)
}
