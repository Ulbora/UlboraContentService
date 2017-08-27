/*
 Copyright (C) 2016 Ulbora Labs Inc. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2016 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published
 by the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
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

//UpdateContentHits in database
func (db *DbConfig) UpdateContentHits(args ...interface{}) bool {
	success := contentDb.UpdateContentHits(args...)
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

//GetContentByClientCategory get a row. Passing in tx allows for transactions
func (db *DbConfig) GetContentByClientCategory(args ...interface{}) *ContentRows {
	var contentRows ContentRows
	rowsPtr := contentDb.GetContentByClientCategory(args...)
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
