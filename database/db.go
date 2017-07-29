package database

import (
	contentDb "UlboraContentService/database/mysqldb"
	"fmt"
	"strconv"
)

//DbConfig db config
type DbConfig struct {
	Host         string
	DbUser       string
	DbPw         string
	DatabaseName string
}

//ContentRow database row
type ContentRow struct {
	Columns []string
	Row     []string
}

//ContentRows array of database rows
type ContentRows struct {
	Columns []string
	Rows    [][]string
}

//ConnectDb to database
func (db *DbConfig) ConnectDb() bool {
	rtn := contentDb.ConnectDb(db.Host, db.DbUser, db.DbPw, db.DatabaseName)
	if rtn == true {
		fmt.Println("db connect")
	}
	return rtn
}

//ConnectionTest of database
func (db *DbConfig) ConnectionTest(args ...interface{}) bool {
	var rtn = false
	rowPtr := contentDb.ConnectionTest(args...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		fmt.Print("Records found during test ")
		fmt.Println(int64Val)
		if err2 != nil {
			fmt.Print(err2)
		}
		if int64Val >= 0 {
			rtn = true
		}
	}
	return rtn
}

//InsertContent in database
func (db *DbConfig) InsertContent(args ...interface{}) (bool, int64) {
	success, insID := contentDb.InsertContent(args...)
	if success == true {
		fmt.Println("inserted record")
	}
	return success, insID
}

//UpdateContent in database
func (db *DbConfig) UpdateContent(args ...interface{}) bool {
	success := contentDb.UpdateContent(args...)
	if success == true {
		fmt.Println("updated record")
	}
	return success
}

//GetContent get a row. Passing in tx allows for transactions
func (db *DbConfig) GetContent(args ...interface{}) *ContentRow {
	var contentRow ContentRow
	rowPtr := contentDb.GetContent(args...)
	if rowPtr != nil {
		contentRow.Columns = rowPtr.Columns
		contentRow.Row = rowPtr.Row
	}
	return &contentRow
}

//GetContentByClient get a row. Passing in tx allows for transactions
func (db *DbConfig) GetContentByClient(args ...interface{}) *ContentRows {
	var contentRows ContentRows
	rowsPtr := contentDb.GetContentByClient(args...)
	if rowsPtr != nil {
		contentRows.Columns = rowsPtr.Columns
		contentRows.Rows = rowsPtr.Rows
	}
	return &contentRows
}

//DeleteContent delete
func (db *DbConfig) DeleteContent(args ...interface{}) bool {
	success := contentDb.DeleteContent(args...)
	return success
}

//CloseDb database connection
func (db *DbConfig) CloseDb() bool {
	rtn := contentDb.CloseDb()
	if rtn == true {
		fmt.Println("db connection closed")
	}
	return rtn
}
