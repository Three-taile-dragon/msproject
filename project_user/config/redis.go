package config

import (
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"test.com/project_user/internal/dao"
)

func (c *Config) ReConnRedis() {
	rdb := redis.NewClient(c.ReadRedisConfig())

	// Enable tracing instrumentation.
	if err := redisotel.InstrumentTracing(rdb); err != nil {
		panic(err)
	}

	// Enable metrics instrumentation.
	if err := redisotel.InstrumentMetrics(rdb); err != nil {
		panic(err)
	}

	rc := &dao.RedisCache{
		Rdb: rdb,
	}
	dao.Rc = rc
}
