package database

//  	"goLang/UlboraContentService/database/mysqldb"

import (
	db "UlboraContentService/database/mysqldb"
	"fmt"
)

//Connect to database
func Connect() bool {
	fmt.Println("db connect")
	return db.ConnectMysql()
}
