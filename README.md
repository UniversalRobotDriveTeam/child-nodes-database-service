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
### func (db *DatabaseAPP) GetComplexConditionsDatas(tableName string, conditions []ComplexCondition, data interface{}) error
当条件复杂(如包含大于、小于等条件时),使用此函数(否则推荐使用`GetData`或`GetDatas`)。传入表的名称、条件和数据结构体指针。返回封装好的错误消息。

条件结构说明：

- 条件(`conditions`)为一个`[]ComplexCondition`结构。其中，`ComplexCondition`为一个结构体，包含`SqlString`和`Condition`两个字段。`SqlString`为SQL语句，`Condition`为SQL语句中的条件。
- 生成条件，推荐调用`GenerateComplexConditions`函数。
- 例如，假设进行以下操作：
```go
condition:=make(map[string]interface{})
condition["user_name = ?"]="test"
condition["age > ?"]=20
condtions,err:=db.GenerateComplexConditions(condition)
db.GetComplexConditionsDatas("users",conditions,&user)
```
等价为：`SELECT * FROM users WHERE user_name = "test" AND age > 20;`

即在`users`表中查找`user_name`为`test`且`age`大于`20`的数据。

---
### func (db *DatabaseAPP) GenerateComplexConditions(conditions map[string]interface{}) []ComplexCondition
生成复杂条件。传入条件(`conditions`)，返回`[]ComplexCondition`结构。

传入结构说明：

- 条件(`conditions`)为一个`map[string]interface{}`结构。其中，`string`为SQL语句(例如:`age > 20`或`user_name = test`),而`interface{}`则是映射条件。