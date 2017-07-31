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

package main

import (
	contentManager "UlboraContentService/manager"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var contentDB contentManager.ContentDB

func main() {
	if os.Getenv("MYSQL_PORT_3306_TCP_ADDR") != "" {
		contentDB.DbConfig.Host = os.Getenv("MYSQL_PORT_3306_TCP_ADDR")
	} else if os.Getenv("DATABASE_HOST") != "" {
		contentDB.DbConfig.Host = os.Getenv("DATABASE_HOST")
	} else {
		contentDB.DbConfig.Host = "localhost:3306"
	}

	if os.Getenv("DATABASE_USER_NAME") != "" {
		contentDB.DbConfig.DbUser = os.Getenv("DATABASE_USER_NAME")
	} else {
		contentDB.DbConfig.DbUser = "admin"
	}

	if os.Getenv("DATABASE_USER_PASSWORD") != "" {
		contentDB.DbConfig.DbPw = os.Getenv("DATABASE_USER_PASSWORD")
	} else {
		contentDB.DbConfig.DbPw = "admin"
	}

	if os.Getenv("DATABASE_NAME") != "" {
		contentDB.DbConfig.DatabaseName = os.Getenv("DATABASE_NAME")
	} else {
		contentDB.DbConfig.DatabaseName = "ulbora_content_service"
	}
	contentDB.ConnectDb()
	defer contentDB.CloseDb()

	fmt.Println("Server running!")
	router := mux.NewRouter()
	router.HandleFunc("/rs/content", handleContentChange).Methods("POST", "PUT")
	router.HandleFunc("/rs/content/{id}/{clientId}", handleContent).Methods("GET", "DELETE")
	router.HandleFunc("/rs/contentList/{clientId}", handleContentList).Methods("GET")
	http.ListenAndServe(":3008", router)
}

func handleContentChange(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	cType := req.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(res, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch req.Method {
		case "POST":
			content := new(contentManager.Content)
			decoder := json.NewDecoder(req.Body)
			error := decoder.Decode(&content)
			if error != nil {
				log.Println(error.Error())
				http.Error(res, error.Error(), http.StatusBadRequest)
			} else if content.Title == "" || content.Text == "" || content.ClientID == 0 {
				http.Error(res, "bad request", http.StatusBadRequest)
			} else {
				content.CreateDate = time.Now()
				fmt.Println(content)
				resOut := contentDB.InsertContent(content)
				fmt.Print("response: ")
				fmt.Println(resOut)
				resJSON, err := json.Marshal(resOut)
				if err != nil {
					log.Println(error.Error())
					http.Error(res, "json output failed", http.StatusInternalServerError)
				}
				res.WriteHeader(http.StatusOK)
				fmt.Fprint(res, string(resJSON))
			}
		case "PUT":
			content := new(contentManager.Content)
			decoder := json.NewDecoder(req.Body)
			error := decoder.Decode(&content)
			if error != nil {
				log.Println(error.Error())
				http.Error(res, error.Error(), http.StatusBadRequest)
			} else if content.Title == "" || content.Text == "" || content.ID == 0 || content.ClientID == 0 {
				http.Error(res, "bad request in update", http.StatusBadRequest)
			} else {
				content.ModifiedDate = time.Now()
				fmt.Println(content)
				resOut := contentDB.UpdateContent(content)
				fmt.Print("response: ")
				fmt.Println(resOut)
				resJSON, err := json.Marshal(resOut)
				if err != nil {
					log.Println(error.Error())
					http.Error(res, "json output failed", http.StatusInternalServerError)
				}
				res.WriteHeader(http.StatusOK)
				fmt.Fprint(res, string(resJSON))
			}
		}
	}
}

func handleContent(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, errID := strconv.ParseInt(vars["id"], 10, 0)
	if errID != nil {
		http.Error(res, "bad request", http.StatusBadRequest)
	}
	clientID, errClient := strconv.ParseInt(vars["clientId"], 10, 0)
	if errClient != nil {
		http.Error(res, "bad request", http.StatusBadRequest)
	}
	fmt.Print("id is: ")
	fmt.Print(id)
	switch req.Method {
	case "GET":
		content := new(contentManager.Content)
		content.ID = id
		content.ClientID = clientID
		resOut := contentDB.GetContent(content)
		fmt.Print("response: ")
		fmt.Println(resOut)
		resJSON, err := json.Marshal(resOut)
		if err != nil {
			log.Println(err.Error())
			http.Error(res, "json output failed", http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res, string(resJSON))
	case "DELETE":
		content := new(contentManager.Content)
		content.ID = id
		content.ClientID = clientID
		resOut := contentDB.DeleteContent(content)
		fmt.Print("response: ")
		fmt.Println(resOut)
		resJSON, err := json.Marshal(resOut)
		if err != nil {
			log.Println(err.Error())
			http.Error(res, "json output failed", http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res, string(resJSON))
	}
}

func handleContentList(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	clientID, errClient := strconv.ParseInt(vars["clientId"], 10, 0)
	if errClient != nil {
		http.Error(res, "bad request", http.StatusBadRequest)
	}
	switch req.Method {
	case "GET":
		content := new(contentManager.Content)
		content.ClientID = clientID
		resOut := contentDB.GetContentByClient(content)
		fmt.Print("response: ")
		fmt.Println(resOut)
		resJSON, err := json.Marshal(resOut)
		if err != nil {
			log.Println(err.Error())
			http.Error(res, "json output failed", http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res, string(resJSON))
	}
}
