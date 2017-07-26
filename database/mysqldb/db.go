package mysqldb

import (
	crud "github.com/Ulbora/go-crud-mysql"
)

//Connect connect to db
func Connect(host, user, pw, dbName string) bool {
	res := crud.InitializeMysql(host, user, pw, dbName)
	return res
}

//GetDb check db
// func GetDb() *sql.DB {
// 	return crud.GetDb()
// }

//Insert insert
func Insert(query string, args ...interface{}) (bool, int64) {
	success, insID := crud.Insert(nil, query, args...)
	return success, insID
}

//Close close connection to db
func Close() bool {
	res := crud.Close()
	return res
}
