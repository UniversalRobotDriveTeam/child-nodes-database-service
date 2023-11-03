package databaseService

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

// Database 数据库暴露出来的读接口
type Database interface {
	ReadItem(item string)
}

// DatabaseAPI 控制数据库的接口
type DatabaseAPI interface {
	// TODO
}

// DatabaseMessageChannel 数据库消息通道
type DatabaseMessageChannel struct {
	// 数据库表地址
	databaseTable string
	// 写入数据库数据通道
	WriteMessageChannel chan DatabaseMessage
	// 从数据库读到数据通道
	ReadMessageChannel chan DatabaseMessage
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
