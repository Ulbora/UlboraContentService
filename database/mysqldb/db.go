package mysqldb

import (
	crud "github.com/Ulbora/go-crud-mysql"
)

//ConnectDb connect to db
func ConnectDb(host, user, pw, dbName string) bool {
	res := crud.InitializeMysql(host, user, pw, dbName)
	return res
}

//InsertContent insert
func InsertContent(args ...interface{}) (bool, int64) {
	success, insID := crud.Insert(nil, InsertContentQuery, args...)
	return success, insID
}

//UpdateContent updates a row. Passing in tx allows for transactions
func UpdateContent(args ...interface{}) bool {
	success := crud.Update(nil, UpdateContentQuery, args...)
	return success
}

//GetContent get a row. Passing in tx allows for transactions
func GetContent(args ...interface{}) *crud.DbRow {
	rowPtr := crud.Get(ContentGetQuery, args...)
	return rowPtr
}

//GetContentByClient get a row. Passing in tx allows for transactions
func GetContentByClient(args ...interface{}) *crud.DbRows {
	rowsPtr := crud.GetList(ContentGetByClientQuery, args...)
	return rowsPtr
}

//DeleteContent delete
func DeleteContent(args ...interface{}) bool {
	success := crud.Delete(nil, DeleteContentQuery, args...)
	return success
}

//CloseDb close connection to db
func CloseDb() bool {
	res := crud.Close()
	return res
}
