package database

import (
	"errors"

	"github.com/238Studio/child-nodes-error-manager/errpack"

	"gorm.io/gorm"
)

// CreateTable 创建表单
// 传入：表单名称，表单结构体
// 传出：错误消息
func (db *DatabaseAPP) CreateTable(tableName string, table interface{}) (err error) {
	defer func() {
		if er := recover(); er != nil {
			//panic错误，定级为fatal
			//返回值赋值
			err = errpack.NewError(errpack.FatalException, errpack.Network, errors.New(er.(string)))
		}
	}()

	err = db.db.Table(tableName).AutoMigrate(table)
	if err != nil {
		return errpack.NewError(errpack.CommonException, errpack.Database, err)
	}
	return nil
}

// AddData 添加数据
// 传入：表单名称，数据(传递指针)
// 传出：错误消息
func (db *DatabaseAPP) AddData(tableName string, data interface{}) (err error) {
	defer func() {
		if er := recover(); er != nil {
			//panic错误，定级为fatal
			//返回值赋值
			err = errpack.NewError(errpack.FatalException, errpack.Network, errors.New(er.(string)))
		}
	}()

	err = db.db.Table(tableName).Create(data).Error
	if err != nil {
		return errpack.NewError(errpack.CommonException, errpack.Database, err)
	}
	return nil
}

// DeleteData 删除数据
// 传入：表单名称，数据(传递指针)
// 传出：错误消息
func (db *DatabaseAPP) DeleteData(tableName string, data interface{}) (err error) {
	defer func() {
		if er := recover(); er != nil {
			//panic错误，定级为fatal
			//返回值赋值
			err = errpack.NewError(errpack.FatalException, errpack.Network, errors.New(er.(string)))
		}
	}()

	err = db.db.Table(tableName).Delete(data).Error
	if err != nil {
		return errpack.NewError(errpack.CommonException, errpack.Database, err)
	}
	return nil
}

// UpdateData 更新数据
// 传入：表单名称，数据(传递指针)
// 传出：错误消息
func (db *DatabaseAPP) UpdateData(tableName string, data interface{}) (err error) {
	defer func() {
		if er := recover(); er != nil {
			//panic错误，定级为fatal
			//返回值赋值
			err = errpack.NewError(errpack.FatalException, errpack.Network, errors.New(er.(string)))
		}
	}()

	err = db.db.Table(tableName).Save(data).Error
	if err != nil {
		return errpack.NewError(errpack.CommonException, errpack.Database, err)
	}
	return nil
}

// GetData 获取一个表单约束字段的首个数据
// 传入：表单名称，数据(传递指针)
// 传出：错误消息
func (db *DatabaseAPP) GetData(tableName string, data interface{}) (err error) {
	defer func() {
		if er := recover(); er != nil {
			//panic错误，定级为fatal
			//返回值赋值
			err = errpack.NewError(errpack.FatalException, errpack.Network, errors.New(er.(string)))
		}
	}()

	err = db.db.Table(tableName).Find(data).Error
	if err != nil {
		return errpack.NewError(errpack.CommonException, errpack.Database, err)
	}
	return nil
}

// GetDatas 获取一个表单约束字段的所有数据
// 传入：表单名称，复合条件，数据(传递指针)
// 传出：错误消息
func (db *DatabaseAPP) GetDatas(tableName string, conditions map[string]interface{}, data interface{}) (err error) {
	defer func() {
		if er := recover(); er != nil {
			//panic错误，定级为fatal
			//返回值赋值
			err = errpack.NewError(errpack.FatalException, errpack.Network, errors.New(er.(string)))
		}
	}()

	err = db.db.Table(tableName).Where(conditions).Find(data).Error
	if err != nil {
		return errpack.NewError(errpack.CommonException, errpack.Database, err)
	}
	return nil
}

// GetComplexConditionsDatas 获取一个表单复杂约束字段的所有数据
// 传入：表单名称，复合条件，数据(传递指针)
// 传出：错误消息
func (db *DatabaseAPP) GetComplexConditionsDatas(tableName string, conditions []ComplexCondition, data interface{}) (err error) {
	defer func() {
		if er := recover(); er != nil {
			//panic错误，定级为fatal
			//返回值赋值
			err = errpack.NewError(errpack.FatalException, errpack.Network, errors.New(er.(string)))
		}
	}()

	table := db.db.Table(tableName)
	//迭代添加约束条件
	for _, condition := range conditions {
		table = table.Where(condition.SqlString, condition.Condition)
	}

	//迭代结束，开始查询操作
	err = table.Find(data).Error
	if err != nil {
		return errpack.NewError(errpack.CommonException, errpack.Database, err)
	}
	return nil
}

// GetDatabase 获取数据库 进行复杂SQL操作
// 传入：无
// 传出：数据库
// FIXME:调用之前需要告知其他开发者。注意。非必要（且不能通过新增接口组合实现）勿用。
func (db *DatabaseAPP) GetDatabase() *gorm.DB {
	return db.db
}
