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
package mysqldb

import (
	crud "github.com/Ulbora/go-crud-mysql"
)

//ConnectDb connect to db
func ConnectDb(host, user, pw, dbName string) bool {
	res := crud.InitializeMysql(host, user, pw, dbName)
	return res
}

//ConnectionTest get a row. Passing in tx allows for transactions
func ConnectionTest(args ...interface{}) *crud.DbRow {
	rowPtr := crud.Get(ConnectionTestQuery, args...)
	return rowPtr
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
