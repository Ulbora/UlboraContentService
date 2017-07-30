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
	"fmt"
	"testing"
	"time"
)

var contentDB ContentDB
var connected bool
var insertID int64
var insertID2 int64

func TestConnectDb(t *testing.T) {
	contentDB.DbConfig.Host = "localhost:3306"
	contentDB.DbConfig.DbUser = "admin"
	contentDB.DbConfig.DbPw = "admin"
	contentDB.DbConfig.DatabaseName = "ulbora_content_service"
	connected = contentDB.ConnectDb()
	if connected != true {
		t.Fail()
	}
}

func TestInsertContent(t *testing.T) {
	var content Content
	content.Title = "test insert in manager"
	content.CreateDate = time.Now()
	content.Hits = 0
	content.MediaAuthorName = "ken"
	content.MediaDesc = "test content"
	content.MediaKeyWords = "test, content, Ulbora"
	content.MediaRobotKeyWorks = "test, content, Ulbora"
	content.Text = "some content text sent from wire"
	content.ClientID = 127

	success, insID := contentDB.InsertContent(content)
	if success == true && insID != -1 {
		insertID = insID
		fmt.Print("new Id: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}

	success2, insID2 := contentDB.InsertContent(content)
	if success2 == true && insID2 != -1 {
		insertID2 = insID2
		fmt.Print("new Id: ")
		fmt.Println(insID2)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestUpdateContent(t *testing.T) {
	var content Content
	content.Title = "test updated in manager"
	content.ModifiedDate = time.Now()
	content.Hits = 5
	content.MediaAuthorName = "ken"
	content.MediaDesc = "test content"
	content.MediaKeyWords = "test, content, Ulbora"
	content.MediaRobotKeyWorks = "test, content, Ulbora"
	content.Text = "some content text sent from wire"
	content.ID = insertID
	content.ClientID = 127
	success := contentDB.UpdateContent(content)
	if success != true {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestUpdateContentHits(t *testing.T) {
	var content Content
	content.ModifiedDate = time.Now()
	content.Hits = 50
	content.ID = insertID2
	content.ClientID = 127
	success := contentDB.UpdateContentHits(content)
	if success != true {
		fmt.Println("database insert failed")
		t.Fail()
	}
}
func TestCloseDb(t *testing.T) {
	success := contentDB.CloseDb()
	if success != true {
		t.Fail()
	}
}
