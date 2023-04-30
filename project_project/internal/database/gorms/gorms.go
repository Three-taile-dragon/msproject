package gorms

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"test.com/project_project/config"
)

var _db *gorm.DB

func init() {
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
