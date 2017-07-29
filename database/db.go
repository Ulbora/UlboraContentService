package database

import (
	contentDb "UlboraContentService/database/mysqldb"
	"fmt"
)

//DbConfig db config
type DbConfig struct {
	Host         string
	DbUser       string
	DbPw         string
	DatabaseName string
}

//ConnectDb to database
func (db *DbConfig) ConnectDb() bool {
	rtn := contentDb.ConnectDb(db.Host, db.DbUser, db.DbPw, db.DatabaseName)
	if rtn == true {
		fmt.Println("db connect")
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

//CloseDb database connection
func (db *DbConfig) CloseDb() bool {
	rtn := contentDb.CloseDb()
	if rtn == true {
		fmt.Println("db connection closed")
	}
	return rtn
}
