package mysqldb

import (
	"fmt"
	"testing"
	"time"
)

var connected bool
var insertID int64

func TestConnect(t *testing.T) {
	connected = Connect("localhost:3306", "admin", "admin", "ulbora_content_service")
	if connected != true {
		fmt.Println("database init failed")
		t.Fail()
	}
}

// func TestGetDb(t *testing.T) {
// 	testDb := GetDb()
// 	if testDb == nil {
// 		fmt.Println("get db failed")
// 		t.Fail()
// 	}
// }

func TestInsert(t *testing.T) {
	//var noTx *sql.Tx
	var q = "INSERT INTO content (title, created_date, text, client_id) VALUES (?, ?, ?, ?)"
	var a []interface{}
	a = append(a, "test insert 2", time.Now(), "some content text", 125)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := Insert(q, a...)
	if success == true && insID != -1 {
		insertID = insID
		fmt.Print("new Id: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}

	// success, insID2 := Insert(noTx, q, a...)
	// if success == true && insID2 != -1 {
	// 	insertID2 = insID2
	// 	fmt.Print("new Id: ")
	// 	fmt.Println(insID2)
	// } else {
	// 	fmt.Println("database insert failed")
	// 	t.Fail()
	// }
}

func TestClose(t *testing.T) {
	if connected == true {
		rtn := Close()
		if rtn != true {
			fmt.Println("database close failed")
			t.Fail()
		}
	}
}
