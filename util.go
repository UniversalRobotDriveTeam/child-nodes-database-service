package database

import (
	"errors"

	"github.com/238Studio/child-nodes-error-manager/errpack"
)

// GenerateComplexConditions 生成复杂条件结构
// 传入:条件键值对
// 传出:条件结构和错误信息
func GenerateComplexConditions(conditions map[string]interface{}) (complexConditions []ComplexCondition, err error) {
	//捕获panic
	defer func() {
		if er := recover(); er != nil {
			//panic错误，定级为fatal
			//返回值赋值
			err = errpack.NewError(errpack.FatalException, errpack.Network, errors.New(er.(string)))
			complexConditions = nil
		}
	}()

	//遍历map，构造条件结构体
	for sqlString, condition := range conditions {
		complexCondition := ComplexCondition{
			SqlString: sqlString,
			Condition: condition,
		}
		complexConditions = append(complexConditions, complexCondition)
	}

	return complexConditions, nil
}
