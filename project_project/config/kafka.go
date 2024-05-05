package config

import (
	"context"
	"go.uber.org/zap"
	"test.com/project_common/kk"
	"test.com/project_project/internal/dao"
	"test.com/project_project/internal/repo"
	"time"
)

var kw *kk.KafkaWriter

func InitKafkaWriter() func() {
	kw = kk.GetWriter("localhost:9092") // TODO 后续参数改为 从配置文件读取
	return kw.Close
}

func SendLog(data []byte) {
	kw.Send(kk.LogData{
		Topic: "msproject_log", // TODO 后续参数改为 从配置文件读取
		Data:  data,
	})
}

// KafkaCache 使用消息队列解决缓存一致性问题
type KafkaCache struct {
	R     *kk.KafkaReader
	cache repo.Cache
}

func NewCacheReader() *KafkaCache {
	// TODO 后续改为配置文件读取
	reader := kk.GetReader([]string{"localhost:9092"}, "cache_group", "msproject_cache")
	return &KafkaCache{
		R:     reader,
		cache: dao.Rc,
	}
}

func (c *KafkaCache) DeleteCache() {
	for {
		message, err := c.R.R.ReadMessage(context.Background())
		if err != nil {
			zap.L().Error("Project DeleteCache ReadMessage err", zap.Error(err))
			continue
		}
		// 删除缓存 key TODO 后续添加不同的判断
		if "task" == string(message.Value) {
			fields, err2 := c.cache.HKeys(context.Background(), "task")
			if err2 != nil {
				zap.L().Error("Project DeleteCache HKeys err", zap.Error(err2))
				continue
			}
			time.Sleep(1 * time.Second) // 延时一秒 删除
			c.cache.Delete(context.Background(), fields)
		}
	}
}

func SendCache(data []byte) {
	kw.Send(kk.LogData{
		Topic: "msproject_cache", // TODO 后续参数改为 从配置文件读取
		Data:  data,
	})
}
