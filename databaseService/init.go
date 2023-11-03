package databaseService

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitSQLiteDatabase 初始化SQLite数据库
// 传入：使用数据库名称，数据库地址
// 传出：数据库对象，错误
func InitSQLiteDatabase(dataBaseName, dataBaseURL string) (*DatabaseAPP, error) {
	databaseApp := new(DatabaseAPP)
	databaseApp.databaseConfig.databaseName = dataBaseName
	databaseApp.databaseConfig.databaseURL = dataBaseURL
	db, err := gorm.Open(sqlite.Open(dataBaseName), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		//TODO:约定的错误处理
		return databaseApp, err
	}
	databaseApp.db = db
	return databaseApp, nil
}
