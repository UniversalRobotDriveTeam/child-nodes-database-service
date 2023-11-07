package databaseService

// CreateTable 创建表单
// 传入：表单名称，表单结构体
// 传出：错误消息
func (db *DatabaseAPP) CreateTable(tableName string, table interface{}) error {
	err := db.db.Table(tableName).AutoMigrate(table)
	return err
}

// AddData 添加数据
// 传入：表单名称，数据(传递指针)
// 传出：错误消息
func (db *DatabaseAPP) AddData(tableName string, data interface{}) error {
	err := db.db.Table(tableName).Create(data).Error
	return err
}

// DeleteData 删除数据
// 传入：表单名称，数据(传递指针)
// 传出：错误消息
func (db *DatabaseAPP) DeleteData(tableName string, data interface{}) error {
	err := db.db.Table(tableName).Delete(data).Error
	return err
}

// UpdateData 更新数据
// 传入：表单名称，数据(传递指针)
// 传出：错误消息
func (db *DatabaseAPP) UpdateData(tableName string, data interface{}) error {
	err := db.db.Table(tableName).Save(data).Error
	return err
}

// GetData 获取数据
// 传入：表单名称，数据(传递指针)
// 传出：错误消息
func (db *DatabaseAPP) GetData(tableName string, data interface{}, conds interface{}) error {
	err := db.db.Table(tableName).Find(data, conds).Error
	return err
}
