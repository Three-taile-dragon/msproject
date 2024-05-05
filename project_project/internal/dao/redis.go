package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var Rc *RedisCache

// RedisCache redis缓存
type RedisCache struct {
	Rdb *redis.Client
}

//func init() {
//	//连接redis 客户端
//	rdb := redis.NewClient(config.C.ReadRedisConfig()) //读取配置文件
//	Rc = &RedisCache{
//		rdb: rdb,
//	}
//}

func (rc *RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) error {
	err := rc.Rdb.Set(ctx, key, value, expire).Err()
	return err
}

func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := rc.Rdb.Get(ctx, key).Result()
	return result, err
}

func (rc *RedisCache) HSet(ctx context.Context, key string, field string, value string) {
	rc.Rdb.HSet(ctx, key, field, value)
	return
}

func (rc *RedisCache) HKeys(ctx context.Context, key string) ([]string, error) {
	result, err := rc.Rdb.HKeys(ctx, key).Result()
	return result, err
}

func (rc *RedisCache) Delete(ctx context.Context, keys []string) {
	rc.Rdb.Del(ctx, keys...)
	return
}
