package database

// GetQueryBody 生成查询体
// 传入:表单名称
// 传出:查询体
func (db *DatabaseAPP) GetQueryBody(tableName string) QueryBody {
	return QueryBody{
		db.db.Table(tableName),
	}
}

// And 添加AND连接条件
// 传入:无
// 返回:无
func (queryBody *QueryBody) And(sqlString string, condition interface{}) {
	queryBody.table = queryBody.table.Where(sqlString, condition)
}

// Or 添加OR连接条件
// 传入:无
// 传出:无
func (queryBody *QueryBody) Or(sqlString string, condition interface{}) {
	queryBody.table = queryBody.table.Or(sqlString, condition)
}

// Not 添加NOT连接条件
// 传入:无
// 传出:无
func (queryBody *QueryBody) Not(sqlString string, condition interface{}) {
	queryBody.table = queryBody.table.Not(sqlString, condition)
}
