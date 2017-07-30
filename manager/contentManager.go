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

package manager

import (
	db "UlboraContentService/database"
	"fmt"
	"time"
)

//Response res
type Response struct {
	Success string
	ID      int64
}

//Content content
type Content struct {
	ID                 int64
	Title              string
	CreateDate         time.Time
	ModifiedDate       time.Time
	Hits               int64
	MediaAuthorName    string
	MediaDesc          string
	MediaKeyWords      string
	MediaRobotKeyWorks string
	Text               string
	ClientID           int64
}

//ContentDB db config
type ContentDB struct {
	DbConfig db.DbConfig
}

//ConnectDb to database
func (db *ContentDB) ConnectDb() bool {
	rtn := db.DbConfig.ConnectDb()
	if rtn == true {
		fmt.Println("db connect")
	}
	return rtn
}

//InsertContent in database
func (db *ContentDB) InsertContent(content Content) (bool, int64) {
	dbConnected := db.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		db.DbConfig.ConnectDb()
	}
	var a []interface{}
	a = append(a, content.Title, content.CreateDate, content.Hits, content.MediaAuthorName, content.MediaDesc, content.MediaKeyWords, content.MediaRobotKeyWorks, content.Text, content.ClientID)
	success, insID := db.DbConfig.InsertContent(a...)
	if success == true {
		fmt.Println("inserted record")
	}
	return success, insID
}

//UpdateContent in database
func (db *ContentDB) UpdateContent(content Content) bool {
	dbConnected := db.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		db.DbConfig.ConnectDb()
	}
	var a []interface{}
	a = append(a, content.Title, content.ModifiedDate, content.Hits, content.MediaAuthorName, content.MediaDesc, content.MediaKeyWords, content.MediaRobotKeyWorks, content.Text, content.ID, content.ClientID)
	success := db.DbConfig.UpdateContent(a...)
	if success == true {
		fmt.Println("update record")
	}
	return success
}

//UpdateContentHits in database
func (db *ContentDB) UpdateContentHits(content Content) bool {
	dbConnected := db.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		db.DbConfig.ConnectDb()
	}
	var a []interface{}
	a = append(a, content.ModifiedDate, content.Hits, content.ID, content.ClientID)
	success := db.DbConfig.UpdateContentHits(a...)
	if success == true {
		fmt.Println("update hits on record")
	}
	return success
}

//CloseDb connection to database
func (db *ContentDB) CloseDb() bool {
	rtn := db.DbConfig.CloseDb()
	if rtn == true {
		fmt.Println("db connect closed")
	}
	return rtn
}
