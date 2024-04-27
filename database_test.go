package database_test

import (
	"testing"

	database "github.com/238Studio/child-nodes-database-service"
)

func TestCreate(t *testing.T) {
	db, err := database.InitSQLiteDatabase("test.db", "")
	if err != nil {
		t.Error(err)
		return
	}

	var TestStruct struct {
		// ID      int
		UserName string
		PassWord string
		Avtar    string
	}
	err = db.CreateTable("testtest1", TestStruct)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestDatabaseAPI(t *testing.T) {
	db, err := database.InitSQLiteDatabase("test.db", "")
	if err != nil {
		t.Error(err)
		return
	}

	var TestStruct struct {
		Id       int
		UserName string
		PassWord string
		Avtar    string
	}
	err = db.CreateTable("testtest", TestStruct)
	if err != nil {
		t.Error(err)
		return
	}

	TestStruct.UserName = "test"
	TestStruct.PassWord = "test"
	TestStruct.Avtar = "test"
	//需要传递指针地址
	err = db.AddData("testtest", &TestStruct)
	if err != nil {
		t.Error(err)
		return
	}

	err = db.GetData("testtest", &TestStruct)
	if err != nil {
		t.Error(err)
		return
	} else {
		t.Log(TestStruct)
	}

	TestStruct.UserName = "testtest"
	TestStruct.PassWord = "testtest"
	err = db.UpdateData("testtest", &TestStruct)
	if err != nil {
		t.Error(err)
		return
	}

	err = db.GetData("testtest", &TestStruct)
	if err != nil {
		t.Error(err)
		return
	} else {
		t.Log(TestStruct)
	}

	err = db.DeleteData("testtest", &TestStruct)
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 获取指定数据
func TestGetData(t *testing.T) {
	db, err := database.InitSQLiteDatabase("test.db", "")
	if err != nil {
		t.Fatal(err)
	}

	var TestStruct struct {
		Id       int
		UserName string
		PassWord string
		Avtar    string
	}
	TestStruct.UserName = "testtest"

	err = db.GetData("testtest", &TestStruct)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(TestStruct)
	type T1 struct {
		Id       int
		UserName string
		PassWord string
		Avtar    string
	}

	test1 := make([]T1, 0)

	condition := make(map[string]interface{})
	condition["user_name"] = "testtest"
	err = db.GetDatas("testtest", condition, &test1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(test1)
}

func TestGetComplexData(t *testing.T) {
	db, err := database.InitSQLiteDatabase("test.db", "")
	if err != nil {
		t.Fatal(err)
	}
	type test struct {
		UserName string
		Id       int
		Age      int
	}
	tests := make([]test, 0)

	cds := db.GetQueryBody("test")
	cds.And("age > ?", 20)
	cds.And("age < ?", 30)
	err = cds.GetComplexConditionsDatas(&tests)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tests)

	tests = make([]test, 0)
	cds.Or("user_name = ?", "testtest")
	err = cds.GetComplexConditionsDatas(&tests)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tests)

	tests = make([]test, 0)
	cds.Not("age = ?", 45)
	err = cds.GetComplexConditionsDatas(&tests)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tests)
}
