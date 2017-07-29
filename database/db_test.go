package database

import (
	"fmt"
	"testing"
	"time"
)

var connected bool
var dbConfig DbConfig
var insertID int64
var insertID2 int64

func TestConnectDb(t *testing.T) {
	//var dbConfig DbConfig
	dbConfig.Host = "localhost:3306"
	dbConfig.DbUser = "admin"
	dbConfig.DbPw = "admin"
	dbConfig.DatabaseName = "ulbora_content_service"
	connected = dbConfig.ConnectDb()
	if connected != true {
		t.Fail()
	}
}

func TestInsertContent(t *testing.T) {
	var a []interface{}
	a = append(a, "test insert 2", time.Now(), "some content text", 126)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := dbConfig.InsertContent(a...)
	if success == true && insID != -1 {
		insertID = insID
		fmt.Print("new Id: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}

	success, insID2 := dbConfig.InsertContent(a...)
	if success == true && insID2 != -1 {
		insertID2 = insID2
		fmt.Print("new Id: ")
		fmt.Println(insID2)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestCloseDb(t *testing.T) {
	if connected == true {
		rtn := dbConfig.CloseDb()
		if rtn != true {
			fmt.Println("database close failed")
			t.Fail()
		}
	}
}
