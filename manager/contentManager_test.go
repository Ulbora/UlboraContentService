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
	content.MetaAuthorName = "ken"
	content.MetaDesc = "test content"
	content.MetaKeyWords = "test, content, Ulbora"
	content.MetaRobotKeyWords = "test, content, Ulbora"
	content.Text = "some content text sent from wire"
	content.ClientID = 127

	res := contentDB.InsertContent(&content)
	if res.Success == true && res.ID != -1 {
		insertID = res.ID
		fmt.Print("new Id: ")
		fmt.Println(res.ID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}

	res2 := contentDB.InsertContent(&content)
	if res2.Success == true && res2.ID != -1 {
		insertID2 = res2.ID
		fmt.Print("new Id: ")
		fmt.Println(res2.ID)
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
	content.MetaAuthorName = "ken"
	content.MetaDesc = "test content"
	content.MetaKeyWords = "test, content, Ulbora"
	content.MetaRobotKeyWords = "test, content, Ulbora"
	content.Text = "some content text sent from wire"
	content.ID = insertID
	content.ClientID = 127
	res := contentDB.UpdateContent(&content)
	if res.Success != true {
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
	res := contentDB.UpdateContentHits(&content)
	if res.Success != true {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestGetContent(t *testing.T) {
	var content Content
	content.ID = insertID
	content.ClientID = 127
	res := contentDB.GetContent(&content)
	fmt.Println("")
	fmt.Print("found content: ")
	fmt.Println(res)
	if res.Hits != 5 {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestGetContentByClient(t *testing.T) {
	var content Content
	content.ClientID = 127
	res := contentDB.GetContentByClient(&content)
	fmt.Println("")
	fmt.Print("found list content: ")
	fmt.Println(res)
	if len(*res) == 0 {
		fmt.Println("database read failed")
		t.Fail()
	} else {
		row1 := (*res)[0]
		if row1.Hits != 5 {
			t.Fail()
		}
		row2 := (*res)[1]
		if row2.Hits != 50 {
			t.Fail()
		}
	}
}

func TestDeleteContent(t *testing.T) {
	var content Content
	content.ID = insertID
	content.ClientID = 127
	res := contentDB.DeleteContent(&content)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}

	content.ID = insertID2
	res2 := contentDB.DeleteContent(&content)
	if res2.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestCloseDb(t *testing.T) {
	success := contentDB.CloseDb()
	if success != true {
		t.Fail()
	}
}
