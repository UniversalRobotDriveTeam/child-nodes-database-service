package database

import (
	_const "github.com/UniversalRobotDriveTeam/child-nodes-assist/const"
	"github.com/UniversalRobotDriveTeam/child-nodes-assist/util"
	"gorm.io/gorm"
)

// CreateTable 创建表单
// 传入：表单名称，表单结构体
// 传出：错误消息
func (db *DatabaseAPP) CreateTable(tableName string, table interface{}) error {
	err := db.db.Table(tableName).AutoMigrate(table)
	if err != nil {
		return util.NewError(_const.CommonException, _const.Database, err)
	}
	return nil
}

// AddData 添加数据
// 传入：表单名称，数据(传递指针)
// 传出：错误消息
func (db *DatabaseAPP) AddData(tableName string, data interface{}) error {
	err := db.db.Table(tableName).Create(data).Error
	if err != nil {
		return util.NewError(_const.CommonException, _const.Database, err)
	}
	return nil
}

// DeleteData 删除数据
// 传入：表单名称，数据(传递指针)
// 传出：错误消息
func (db *DatabaseAPP) DeleteData(tableName string, data interface{}) error {
	err := db.db.Table(tableName).Delete(data).Error
	if err != nil {
		return util.NewError(_const.CommonException, _const.Database, err)
	}
	return nil
}

// UpdateData 更新数据
// 传入：表单名称，数据(传递指针)
// 传出：错误消息
func (db *DatabaseAPP) UpdateData(tableName string, data interface{}) error {
	err := db.db.Table(tableName).Save(data).Error
	if err != nil {
		return util.NewError(_const.CommonException, _const.Database, err)
	}
	return nil
}

// GetData 获取一个表单全部数据
// 传入：表单名称，数据(传递指针)
// 传出：错误消息
func (db *DatabaseAPP) GetData(tableName string, data interface{}) error {
	err := db.db.Table(tableName).Find(data).Error
	if err != nil {
		return util.NewError(_const.CommonException, _const.Database, err)
	}
	return nil
}

// GetDatabase 获取数据库 进行复杂SQL操作
// 传入：无
// 传出：数据库
// FIXME:这么做是否合适我先蒙古。
func (db *DatabaseAPP) GetDatabase() *gorm.DB {
	return db.db
}
