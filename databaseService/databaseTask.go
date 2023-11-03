package databaseService

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Start 启动模块服务
// 传入：启动参数
// 传出：无
func (app *DatabaseAPP) Start() {
	db, err := gorm.Open(sqlite.Open(app.databaseConfig.databaseName), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return
	}
	//todo:err
	app.db = db
}

// Stop 中止模块服务
// 传入：无
// 传出：无
func (app *DatabaseAPP) Stop() {
	return
}

// GetApp 获取App
// 传入：无
// 传出：该模块App的指针
func (app *DatabaseAPP) GetApp() *interface{} {
	var value interface{}
	value = app
	return &value
}

// IsAlive 是否在服务
// 传入：无
// 传出：是否在服务
func (app *DatabaseAPP) IsAlive() bool {
	err := app.UpdateData("test", "")
	if err != nil {
		return false
		//todo
	}
	if err != nil {
		return false
	}
	return true
}
