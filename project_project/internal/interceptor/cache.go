package interceptor

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"strings"
	"test.com/project_common/encrypts"
	"test.com/project_grpc/project"
	"test.com/project_grpc/task"
	"test.com/project_project/internal/dao"
	"test.com/project_project/internal/repo"
	"time"
)

// CacheInterceptor 除了缓存拦截器 实现日志拦截器 打印参数内容值 请求的时间 等等
type CacheInterceptor struct {
	cache    repo.Cache
	cacheMap map[string]CacheRespOption
}

type CacheRespOption struct {
	path   string
	typ    any
	expire time.Duration
}

func New() *CacheInterceptor {

	//cacheMap := make(map[string]any)
	//cacheMap["/task.service.v1.TaskService/TaskList"] = &task.TaskListResponse{}
	//
	//缓存接口列表
	cacheMap := map[string]CacheRespOption{
		"/project.ProjectService/Index": {
			path:   "/project.service.v1.ProjectService/Index",
			typ:    &project.IndexResponse{},
			expire: 1 * time.Hour,
		},
		"/task.service.v1.TaskService/TaskList": {
			path:   "/task.service.v1.TaskService/TaskList",
			typ:    &task.TaskListResponse{},
			expire: 1 * time.Hour,
		},
	}
	return &CacheInterceptor{cache: dao.Rc, cacheMap: cacheMap}
}

func (c *CacheInterceptor) Cache() grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		c = New()
		respOption, ok := c.cacheMap[info.FullMethod]
		if !ok { //路径不在缓存列表内
			return handler(ctx, req)
		}
		//先查询是否有缓存 有直接返回，无 先请求后存入缓存
		con, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		//redis key 由 req 进行 MD5加密得到
		reqJson, _ := json.Marshal(req)
		cacheKey := encrypts.Md5(string(reqJson))
		respJson, _ := c.cache.Get(con, info.FullMethod+"::"+cacheKey)
		if respJson != "" {
			err := json.Unmarshal([]byte(respJson), respOption.typ)
			return respOption.typ, err
		}

		resp, err = handler(ctx, req)
		respJson2, _ := json.Marshal(resp)
		err = c.cache.Put(con, info.FullMethod+"::"+cacheKey, string(respJson2), respOption.expire)

		// hash key task field redisKey
		// 设置缓存 key TODO 后续添加不同的判断
		if strings.HasPrefix(info.FullMethod, "/task") {
			c.cache.HSet(con, "task", info.FullMethod+"::"+cacheKey, "")
		}

		return resp, err
	})
}

func (c *CacheInterceptor) CacheInterceptor() func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		c = New()
		respOption, ok := c.cacheMap[info.FullMethod]
		if !ok { //路径不在缓存列表内
			return handler(ctx, req)
		}
		//先查询是否有缓存 有直接返回，无 先请求后存入缓存
		con, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		//redis key 由 req 进行 MD5加密得到
		reqJson, _ := json.Marshal(req)
		cacheKey := encrypts.Md5(string(reqJson))
		respJson, _ := c.cache.Get(con, info.FullMethod+"::"+cacheKey)
		if respJson != "" {
			err := json.Unmarshal([]byte(respJson), respOption.typ)
			return respOption.typ, err
		}

		resp, err = handler(ctx, req)
		respJson2, _ := json.Marshal(resp)
		err = c.cache.Put(con, info.FullMethod+"::"+cacheKey, string(respJson2), respOption.expire)

		// hash key task field redisKey
		// 设置缓存 key TODO 后续添加不同的判断
		if strings.HasPrefix(info.FullMethod, "/task") {
			c.cache.HSet(con, "task", info.FullMethod+"::"+cacheKey, "")
		}

		return
	}
}
