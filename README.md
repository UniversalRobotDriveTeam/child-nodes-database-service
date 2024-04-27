# child-nodes-database-service

子节点数据库服务。

## API说明

### func InitSQLiteDatabase(dataBaseName, dataBaseURL string) (*DatabaseAPP, error)
创造一个SQLite数据库APP。传入数据库的名称和地址（考虑到远程调用的可能性）。返回数据库APP与封装好的错误消息。

---
### func (db *DatabaseAPP) CreateTable(tableName string, table interface{}) error
创建一个表。传入表的名称和表的结构体指针。返回封装好的错误消息。

---
### func (db *DatabaseAPP) AddData(tableName string, data interface{}) error
向表中添加数据。传入表的名称和数据结构体指针。返回封装好的错误消息。

---
### func (db *DatabaseAPP) DeleteData(tableName string, data interface{}) error
从表中删除数据。传入表的名称和数据结构体指针。返回封装好的错误消息。

---
### func (db *DatabaseAPP) UpdateData(tableName string, data interface{}) error
更新表中的数据。传入表的名称和数据结构体指针。返回封装好的错误消息。

---
### func (db *DatabaseAPP) DeleteData(tableName string, data interface{}) error
从表中删除数据。传入表的名称和数据结构体指针。返回封装好的错误消息。

---
### func (db *DatabaseAPP) GetData(tableName string, data interface{}) error
从表中获取首个数据。传入表的名称和数据结构体指针。返回封装好的错误消息。

---
### func (db *DatabaseAPP) GetDatas(tableName string, conditions map[string]interface{}, data interface{}) error
从表中获取多个数据。传入表的名称、条件和数据结构体指针。返回封装好的错误消息。

条件结构说明：

- 条件(`conditions`)为一个`map[string]interface{}`结构。其中，`string`为SQL数据库中的字段名称(注意：非`gorm`规范的驼峰命名，而是蛇形命名),而`interface{}`则是映射条件。
- 例如，假设进行以下操作：
```go
conditions:=make(map[string]interface{})
conditions["user_name"]="test"
conditions["age"]=20
db.GetDatas("users",conditions,&user)
```
等价为：`SELECT * FROM users WHERE user_name = "test" AND age = 20;`

即在`users`表中查找`user_name`为`test`且`age`为`20`的数据。

---

### func (db *DatabaseAPP) GetQueryBody(tableName string) QueryBody
获取查询体。传入表的名称。返回查询体。

---
### func (queryBody *QueryBody) And(sqlString string, condition interface{})
添加AND条件。传入SQL语句和条件。更新查询体查询信息。

---
### func (queryBody *QueryBody) Or(sqlString string, condition interface{})
添加OR条件。传入SQL语句和条件。更新查询体查询信息。

---
### func (queryBody *QueryBody) Not(sqlString string, condition interface{})
添加NOT条件。传入SQL语句和条件。更新查询体查询信息。

---
### func (queryBody *QueryBody) GetComplexConditionsDatas(data interface{}) (err error)
执行复杂条件查询。传入数据结构体指针。返回封装好的错误消息。