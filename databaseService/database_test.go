package databaseService_test

import (
	"github.com/UniversalRobotDriveTeam/child-nodes-basic/robotBasicAPI/databaseService"
	"testing"
)

func TestCreate(t *testing.T) {
	db, err := databaseService.InitSQLiteDatabase("test.db", "")
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
	db, err := databaseService.InitSQLiteDatabase("test.db", "")
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
