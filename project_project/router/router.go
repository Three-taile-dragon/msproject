package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
	"test.com/project_common/discovery"
	"test.com/project_common/logs"
	"test.com/project_grpc/account"
	"test.com/project_grpc/department"
	project_service "test.com/project_grpc/project"
	"test.com/project_grpc/task"
	"test.com/project_project/config"
	"test.com/project_project/internal/interceptor"
	"test.com/project_project/internal/rpc"
	account_service_v1 "test.com/project_project/pkg/service/account.service.v1"
	department_service_v1 "test.com/project_project/pkg/service/department.service.v1"
	project_service_v1 "test.com/project_project/pkg/service/project.service.v1"
	task_service_v1 "test.com/project_project/pkg/service/task.service.v1"
)

// Router 接口
type Router interface {
	Router(r *gin.Engine)
}

//type RegisterRouter struct {
//}

//func New() *RegisterRouter {
//	return &RegisterRouter{}
//}

//func (*RegisterRouter) Route(ro Router, r *gin.Engine) {
//	ro.Router(r)
//}

// routers 使用路由列表
// 后续添加路由，不需要再对InitRouter函数进行改动
var routers []Router

// InitRouter 初始化路由
func InitRouter(r *gin.Engine) {
	//reg := New()
	//路由注册
	//reg.Route(&user.RouteruUser{}, r)
	for _, reg := range routers {
		reg.Router(r)
	}

}

// Register 添加到路由列表中去
func Register(ro ...Router) {
	routers = append(routers, ro...)
}

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

// RegisterGrpc 注册grpc服务
func RegisterGrpc() *grpc.Server {
	c := gRPCConfig{
		Addr: config.C.GC.Addr,
		RegisterFunc: func(g *grpc.Server) {
			project_service.RegisterProjectServiceServer(g, project_service_v1.New())
			task.RegisterTaskServiceServer(g, task_service_v1.New())
			account.RegisterAccountServiceServer(g, account_service_v1.New())
			department.RegisterDepartmentServiceServer(g, department_service_v1.New())
		}}
	// grpc 拦截器	自定义统一缓存
	cacheInterceptor := interceptor.New()
	s := grpc.NewServer(cacheInterceptor.Cache()) //启动grpc服务
	c.RegisterFunc(s)                             //注册grpc登陆模块
	lis, err := net.Listen("tcp", config.C.GC.Addr)
	if err != nil {
		log.Println("cannot listen")
	}
	//放到协程里面 防止阻塞主进程main
	go func() {
		log.Printf("grpc server started as %s \n", c.Addr)
		err = s.Serve(lis)
		if err != nil {
			log.Println("server started error", err)
			return
		}
	}()
	return s
}

func RegisterEtcdServer() {
	etcdRegister := discovery.NewResolver(config.C.EC.Addrs, logs.LG)
	resolver.Register(etcdRegister)
	info := discovery.Server{
		Name:    config.C.GC.Name,
		Addr:    config.C.GC.Addr,
		Version: config.C.GC.Version,
		Weight:  config.C.GC.Weight,
	}
	r := discovery.NewRegister(config.C.EC.Addrs, logs.LG)
	_, err := r.Register(info, 2)
	if err != nil {
		log.Fatalln(err)
	}
}

func InitUserRpc() {
	rpc.InitRpcUserClient()
}
