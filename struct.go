package database

import (
	"gorm.io/gorm"
)

// DatabaseAPP 数据库调用结构
type DatabaseAPP struct {
	// 数据库配置
	databaseConfig databaseConfig
	// 数据库引擎
	db *gorm.DB
	// 调用数据库的方法
	DatabaseChannel *chan map[string]interface{}
}

// 数据库应用的配置文件
type databaseConfig struct {
	// 数据库名
	databaseName string
	// 数据库位置
	databaseURL string
	// 数据库用户名
	databaseUserName string
	// 数据库用户密码
	databaseUserPassword string
	// TODO

}

// DatabaseMessage 数据库写入对象
type DatabaseMessage struct {
	// 数据项目ID
	id string
	// 数据项目主键
	key string
	// 数据
	strData   map[string]string
	intData   map[string]int64
	floatData map[string]float64
}
