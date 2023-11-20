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
}

// ComplexCondition 复杂数据查询约束条件
type ComplexCondition struct {
	SqlString string
	Condition interface{}
}
