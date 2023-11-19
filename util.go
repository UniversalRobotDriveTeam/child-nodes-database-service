package database

import (
	"errors"

	_const "github.com/238Studio/child-nodes-assist/const"
	"github.com/238Studio/child-nodes-assist/util"
)

func GenerateComplexConditions(conditions map[string]interface{}) ([]ComplexCondition, error) {
	//捕获panic
	defer func() {
		if err := recover(); err != nil {
			//panic错误，定级为fatal
			//返回值赋值
			err = util.NewError(_const.FatalException, _const.Network, errors.New(err.(string)))
		}
	}()

	var complexConditions []ComplexCondition

	for sqlString, condition := range conditions {
		complexCondition := ComplexCondition{
			SqlString: sqlString,
			Condition: condition,
		}
		complexConditions = append(complexConditions, complexCondition)
	}

	return complexConditions, nil
}
