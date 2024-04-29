package config

import (
	"bytes"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
	"test.com/project_common/logs"
)

var C = InitConfig()

type Config struct {
	viper *viper.Viper
	SC    *ServerConfig
	GC    *GrpcConfig
	EC    *EtcdConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

type GrpcConfig struct {
	Name string
	Addr string
}
type EtcdConfig struct {
	Addrs []string
}

// InitConfig 读取配置文件
func InitConfig() *Config {
	conf := &Config{viper: viper.New()}
	// 先从 nacos 读取配置， 如果读取不到 再从本地读取
	nacosClient := InitNacosClient()
	configYaml, err2 := nacosClient.configClient.GetConfig(vo.ConfigParam{
		DataId: "config.yaml",
		Group:  nacosClient.group,
	})
	if err2 != nil {
		log.Fatalln(err2)
	}
	conf.viper.SetConfigType("yaml") //viper配置文件后缀
	if configYaml != "" {
		err := conf.viper.ReadConfig(bytes.NewBuffer([]byte(configYaml)))
		if err != nil {
			log.Fatalln(err) //报错就退出
		}
		//log.Printf("load nacos config %s \n", configYaml)
	} else {
		workDir, _ := os.Getwd()
		conf.viper.SetConfigName("config") //viper 配置文件名次

		//conf.viper.AddConfigPath("/etc/ms_project/api") //配置文件目录 可添加多个 按照代码顺序读取
		conf.viper.AddConfigPath(workDir + "/config")
		//读入配置
		err := conf.viper.ReadInConfig()
		if err != nil {
			zap.L().Error("config read wrong, err: " + err.Error())
			log.Fatalln(err) //报错就退出
		}
	}

	conf.ReadServerConfig()
	conf.InitZapLog()
	conf.ReadEtcdConfig()
	return conf
}

// ReadServerConfig 读取服务器地址配置
func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name") //读取配置
	sc.Addr = c.viper.GetString("server.addr")
	c.SC = sc
}

// ReadEtcdConfig 读取etcd地址配置
func (c *Config) ReadEtcdConfig() {
	ec := &EtcdConfig{}
	var addrs []string
	err := c.viper.UnmarshalKey("etcd.addrs", &addrs)
	if err != nil {
		log.Fatalf("etcd config read wrong, err :%v \n", err)
	}
	ec.Addrs = addrs
	c.EC = ec
}
func (c *Config) InitZapLog() {
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("maxSize"),
		MaxAge:        c.viper.GetInt("maxAge"),
		MaxBackups:    c.viper.GetInt("maxBackups"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}
