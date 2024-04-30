package config

import "test.com/project_common/kk"

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
