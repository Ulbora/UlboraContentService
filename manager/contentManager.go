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
	"strconv"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

//Response res
type Response struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
}

//Content content
type Content struct {
	ID                int64     `json:"id"`
	Title             string    `json:"title"`
	CreateDate        time.Time `json:"createDate"`
	ModifiedDate      time.Time `json:"modifiedDate"`
	Hits              int64     `json:"hits"`
	MetaAuthorName    string    `json:"metaAuthorName"`
	MetaDesc          string    `json:"metaDesc"`
	MetaKeyWords      string    `json:"metaKeyWords"`
	MetaRobotKeyWords string    `json:"metaRobotKeyWords"`
	Text              string    `json:"text"`
	ClientID          int64     `json:"clientId"`
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
func (db *ContentDB) InsertContent(content *Content) *Response {
	var rtn Response
	dbConnected := db.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		db.DbConfig.ConnectDb()
	}
	var a []interface{}
	a = append(a, content.Title, content.CreateDate, content.Hits, content.MetaAuthorName, content.MetaDesc, content.MetaKeyWords, content.MetaRobotKeyWords, content.Text, content.ClientID)
	success, insID := db.DbConfig.InsertContent(a...)
	if success == true {
		fmt.Println("inserted record")
	}
	rtn.ID = insID
	rtn.Success = success
	return &rtn
}

//UpdateContent in database
func (db *ContentDB) UpdateContent(content *Content) *Response {
	var rtn Response
	dbConnected := db.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		db.DbConfig.ConnectDb()
	}
	var a []interface{}
	a = append(a, content.Title, content.ModifiedDate, content.Hits, content.MetaAuthorName, content.MetaDesc, content.MetaKeyWords, content.MetaRobotKeyWords, content.Text, content.ID, content.ClientID)
	success := db.DbConfig.UpdateContent(a...)
	if success == true {
		fmt.Println("update record")
	}
	rtn.ID = content.ID
	rtn.Success = success
	return &rtn
}

//UpdateContentHits in database
func (db *ContentDB) UpdateContentHits(content *Content) *Response {
	var rtn Response
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
	rtn.ID = content.ID
	rtn.Success = success
	return &rtn
}

//GetContent content from database
func (db *ContentDB) GetContent(content *Content) *Content {
	var a []interface{}
	a = append(a, content.ID, content.ClientID)
	var rtn *Content
	rowPtr := db.DbConfig.GetContent(a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		rtn = parseContentRow(&foundRow)
	}
	return rtn
}

//GetContentByClient content by Client
func (db *ContentDB) GetContentByClient(content *Content) *[]Content {
	var rtn []Content
	var a []interface{}
	a = append(a, content.ClientID)
	rowsPtr := db.DbConfig.GetContentByClient(a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := parseContentRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteContent in database
func (db *ContentDB) DeleteContent(content *Content) *Response {
	var rtn Response
	dbConnected := db.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		db.DbConfig.ConnectDb()
	}
	var a []interface{}
	a = append(a, content.ID, content.ClientID)
	success := db.DbConfig.DeleteContent(a...)
	if success == true {
		fmt.Println("deleted record")
	}
	rtn.ID = content.ID
	rtn.Success = success
	return &rtn
}

//CloseDb connection to database
func (db *ContentDB) CloseDb() bool {
	rtn := db.DbConfig.CloseDb()
	if rtn == true {
		fmt.Println("db connect closed")
	}
	return rtn
}

func parseContentRow(foundRow *[]string) *Content {
	var rtn Content
	id, errID := strconv.ParseInt((*foundRow)[0], 10, 0)
	if errID != nil {
		fmt.Print(errID)
	} else {
		rtn.ID = id
	}
	rtn.Title = (*foundRow)[1]
	cTime, errCtime := time.Parse(timeFormat, (*foundRow)[2])
	if errCtime != nil {
		fmt.Print(errCtime)
	} else {
		rtn.CreateDate = cTime
	}
	mTime, errMtime := time.Parse(timeFormat, (*foundRow)[3])
	if errMtime != nil {
		fmt.Print(errMtime)
	} else {
		rtn.ModifiedDate = mTime
	}
	hits, errHits := strconv.ParseInt((*foundRow)[4], 10, 0)
	if errHits != nil {
		fmt.Print(errHits)
	} else {
		rtn.Hits = hits
	}
	rtn.MetaAuthorName = (*foundRow)[5]
	rtn.MetaDesc = (*foundRow)[6]
	rtn.MetaKeyWords = (*foundRow)[7]
	rtn.MetaRobotKeyWords = (*foundRow)[8]
	rtn.Text = (*foundRow)[9]
	clientID, errClient := strconv.ParseInt((*foundRow)[10], 10, 0)
	if errClient != nil {
		fmt.Print(errClient)
	} else {
		rtn.ClientID = clientID
	}
	return &rtn
}
