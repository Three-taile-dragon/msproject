package interceptor

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"test.com/project_common/encrypts"
	"test.com/project_grpc/project"
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
	//缓存接口列表
	cacheMap := map[string]CacheRespOption{
		"/project.ProjectService/Index": {
			path:   "/project.service.v1.ProjectService/Index",
			typ:    &project.IndexResponse{},
			expire: 1 * time.Hour,
		},
		"/project.ProjectService/FindProjectByMemId": {
			path:   "/project.service.v1.ProjectService/FindProjectByMemId",
			typ:    &project.MyProjectResponse{},
			expire: 10 * time.Minute,
		},
	}
	return &CacheInterceptor{cache: dao.Rc, cacheMap: cacheMap}
}

func (c *CacheInterceptor) Cache() grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
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
		return resp, err
	})
}
