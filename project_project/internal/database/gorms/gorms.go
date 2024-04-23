package gorms

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"test.com/project_project/config"
)

var _db *gorm.DB

func init() {
	// 判断是否开启 读写分离
	if config.C.DC.Separation {
		// 开启读写分离
		username := config.C.DC.Master.Username //账号
		password := config.C.DC.Master.Password //密码
		host := config.C.DC.Master.Host         //数据库地址，可以是Ip或者域名
		port := config.C.DC.Master.Port         //数据库端口
		Dbname := config.C.DC.Master.Db         //数据库名
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
		var err error
		_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			zap.L().Error("数据库连接失败", zap.Error(err))
			panic("连接数据库失败, error=" + err.Error())
		}
		// slave
		var replicas []gorm.Dialector
		for _, v := range config.C.DC.Slave {
			username := v.Username //账号
			password := v.Password //密码
			host := v.Host         //数据库地址，可以是Ip或者域名
			port := v.Port         //数据库端口
			Dbname := v.Db         //数据库名
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
			cfg := mysql.Config{
				DSN: dsn,
			}
			replicas = append(replicas, mysql.New(cfg))
		}
		// 加载读写分离
		_ = _db.Use(dbresolver.Register(dbresolver.Config{
			// 主库
			Sources: []gorm.Dialector{mysql.New(mysql.Config{
				DSN: dsn,
			})},
			// 从库
			Replicas: replicas,
			Policy:   dbresolver.RandomPolicy{},
		}).SetMaxOpenConns(200).SetMaxIdleConns(10)) // 连接池配置
	} else {
		//配置MySQL连接参数
		username := config.C.MC.Username //账号
		password := config.C.MC.Password //密码
		host := config.C.MC.Host         //数据库地址，可以是Ip或者域名
		port := config.C.MC.Port         //数据库端口
		Dbname := config.C.MC.Db         //数据库名
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
		var err error
		_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			zap.L().Error("数据库连接失败", zap.Error(err))
			panic("连接数据库失败, error=" + err.Error())
		}
	}

}
func GetDB() *gorm.DB {
	return _db
}

type GormConn struct {
	db *gorm.DB
	tx *gorm.DB
}

func (g *GormConn) Begin() {
	//g.tx = GetDB().Begin()
	tx := GetDB().Begin()
	if tx.Error != nil {
		zap.L().Error("GormConn.Begin error", zap.Error(tx.Error))
		return
	}
	g.tx = tx
}

func New() *GormConn {
	return &GormConn{db: GetDB(), tx: nil}
}

func NewTran() *GormConn {
	//return &GormConn{db: GetDB(), tx: GetDB().Begin()}
	tx := GetDB().Begin()
	if tx.Error != nil {
		zap.L().Error("GormConn.NewTran error", zap.Error(tx.Error))
		return nil
	}
	return &GormConn{db: GetDB(), tx: tx}
}
func (g *GormConn) Session(ctx context.Context) *gorm.DB {
	return g.db.Session(&gorm.Session{Context: ctx})
}

// 事务

func (g *GormConn) Rollback() {
	g.tx.Rollback()
}
func (g *GormConn) Commit() {
	g.tx.Commit()
}

func (g *GormConn) Tx(ctx context.Context) *gorm.DB {
	return g.tx.WithContext(ctx)
}
