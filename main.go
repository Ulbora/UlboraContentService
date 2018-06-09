/*
 Copyright (C) 2016 Ulbora Labs LLC. (www.ulboralabs.com)
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
	"strings"
	"time"

	uoauth "github.com/Ulbora/go-ulbora-oauth2"
	"github.com/gorilla/mux"
)

var contentDB contentManager.ContentDB

type authHeader struct {
	token    string
	clientID int64
	userID   string
	hashed   bool
}

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

	fmt.Println("Server running on port 3008!")
	router := mux.NewRouter()
	router.HandleFunc("/rs/content/add", handleContentChange).Methods("POST")
	router.HandleFunc("/rs/content/update", handleContentChange).Methods("PUT")
	router.HandleFunc("/rs/content/hits", handleContentHits).Methods("PUT")
	router.HandleFunc("/rs/content/get/{id}/{clientId}", handleContent).Methods("GET")
	router.HandleFunc("/rs/content/list/{clientId}", handleContentList).Methods("GET")
	router.HandleFunc("/rs/content/list/{clientId}/{category}", handleContentListCategory).Methods("GET")
	router.HandleFunc("/rs/content/delete/{id}", handleContent).Methods("DELETE")
	http.ListenAndServe(":3008", router)
}

func handleContentChange(res http.ResponseWriter, req *http.Request) {
	auth := getAuth(req)
	me := new(uoauth.Claim)
	me.Role = "admin"
	me.Scope = "write"
	res.Header().Set("Content-Type", "application/json")
	cType := req.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(res, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch req.Method {
		case "POST":
			me.URI = "/ulbora/rs/content/add"
			valid := auth.Authorize(me)
			if valid != true {
				res.WriteHeader(http.StatusUnauthorized)
			} else {
				content := new(contentManager.Content)
				decoder := json.NewDecoder(req.Body)
				error := decoder.Decode(&content)
				content.ClientID = auth.ClientID
				if error != nil {
					log.Println(error.Error())
					http.Error(res, error.Error(), http.StatusBadRequest)
				} else if content.Title == "" || content.Text == "" || content.ClientID == 0 {
					http.Error(res, "bad request", http.StatusBadRequest)
				} else {
					content.CreateDate = time.Now()
					//fmt.Println(content)
					resOut := contentDB.InsertContent(content)
					//fmt.Print("response: ")
					//fmt.Println(resOut)
					resJSON, err := json.Marshal(resOut)
					if err != nil {
						log.Println(error.Error())
						http.Error(res, "json output failed", http.StatusInternalServerError)
					}
					res.WriteHeader(http.StatusOK)
					fmt.Fprint(res, string(resJSON))
				}
			}
		case "PUT":
			me.URI = "/ulbora/rs/content/update"
			valid := auth.Authorize(me)
			if valid != true {
				res.WriteHeader(http.StatusUnauthorized)
			} else {
				content := new(contentManager.Content)
				decoder := json.NewDecoder(req.Body)
				error := decoder.Decode(&content)
				content.ClientID = auth.ClientID
				if error != nil {
					log.Println(error.Error())
					http.Error(res, error.Error(), http.StatusBadRequest)
				} else if content.Title == "" || content.Text == "" || content.ID == 0 || content.ClientID == 0 {
					http.Error(res, "bad request in update", http.StatusBadRequest)
				} else {
					content.ModifiedDate = time.Now()
					//fmt.Println(content)
					resOut := contentDB.UpdateContent(content)
					//fmt.Print("response: ")
					//fmt.Println(resOut)
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
}

func handleContentHits(res http.ResponseWriter, req *http.Request) {
	auth := getAuth(req)
	me := new(uoauth.Claim)
	me.Role = "admin"
	me.Scope = "write"
	res.Header().Set("Content-Type", "application/json")
	cType := req.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(res, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch req.Method {
		case "PUT":
			me.URI = "/ulbora/rs/content/hits"
			valid := auth.Authorize(me)
			if valid != true {
				res.WriteHeader(http.StatusUnauthorized)
			} else {
				content := new(contentManager.Content)
				decoder := json.NewDecoder(req.Body)
				error := decoder.Decode(&content)
				content.ClientID = auth.ClientID
				if error != nil {
					log.Println(error.Error())
					http.Error(res, error.Error(), http.StatusBadRequest)
				} else if content.ID == 0 || content.ClientID == 0 {
					http.Error(res, "bad request in update", http.StatusBadRequest)
				} else {
					content.ModifiedDate = time.Now()
					//fmt.Println(content)
					resOut := contentDB.UpdateContentHits(content)
					//fmt.Print("response: ")
					//fmt.Println(resOut)
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
}

func handleContent(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, errID := strconv.ParseInt(vars["id"], 10, 0)
	if errID != nil {
		http.Error(res, "bad request", http.StatusBadRequest)
	}
	//fmt.Print("id is: ")
	//fmt.Println(id)
	switch req.Method {
	case "GET":
		clientID, errClient := strconv.ParseInt(vars["clientId"], 10, 0)
		if errClient != nil {
			http.Error(res, "bad request", http.StatusBadRequest)
		}
		content := new(contentManager.Content)
		content.ID = id
		content.ClientID = clientID
		resOut := contentDB.GetContent(content)
		//fmt.Print("response: ")
		//fmt.Println(resOut)
		resJSON, err := json.Marshal(resOut)
		if err != nil {
			log.Println(err.Error())
			http.Error(res, "json output failed", http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res, string(resJSON))
	case "DELETE":
		auth := getAuth(req)
		me := new(uoauth.Claim)
		me.Role = "admin"
		me.Scope = "write"
		me.URI = "/ulbora/rs/content/delete"
		valid := auth.Authorize(me)
		if valid != true {
			res.WriteHeader(http.StatusUnauthorized)
		} else {
			content := new(contentManager.Content)
			content.ID = id
			content.ClientID = auth.ClientID
			resOut := contentDB.DeleteContent(content)
			//fmt.Print("response: ")
			//fmt.Println(resOut)
			resJSON, err := json.Marshal(resOut)
			if err != nil {
				log.Println(err.Error())
				http.Error(res, "json output failed", http.StatusInternalServerError)
			}
			res.WriteHeader(http.StatusOK)
			fmt.Fprint(res, string(resJSON))
		}
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
		//fmt.Print("response: ")
		//fmt.Println(resOut)

		resJSON, err := json.Marshal(resOut)
		//fmt.Print("response json: ")
		//fmt.Println(string(resJSON))
		if err != nil {
			log.Println(err.Error())
			http.Error(res, "json output failed", http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusOK)
		if string(resJSON) == "null" {
			fmt.Fprint(res, "[]")
		} else {
			fmt.Fprint(res, string(resJSON))
		}

	}
}

func handleContentListCategory(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	clientID, errClient := strconv.ParseInt(vars["clientId"], 10, 0)
	if errClient != nil {
		http.Error(res, "bad request", http.StatusBadRequest)
	}
	category := vars["category"]
	switch req.Method {
	case "GET":
		content := new(contentManager.Content)
		content.ClientID = clientID
		content.Category = category
		resOut := contentDB.GetContentByClientCategory(content)
		//fmt.Print("response: ")
		//fmt.Println(resOut)

		resJSON, err := json.Marshal(resOut)
		//fmt.Print("response json: ")
		//fmt.Println(string(resJSON))
		if err != nil {
			log.Println(err.Error())
			http.Error(res, "json output failed", http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusOK)
		if string(resJSON) == "null" {
			fmt.Fprint(res, "[]")
		} else {
			fmt.Fprint(res, string(resJSON))
		}

	}
}

func getHeaders(req *http.Request) *authHeader {
	var rtn = new(authHeader)
	authHeader := req.Header.Get("Authorization")
	tokenArray := strings.Split(authHeader, " ")
	if len(tokenArray) == 2 {
		rtn.token = tokenArray[1]
		//fmt.Println(rtn.token)
	}
	userIDHeader := req.Header.Get("userId")
	rtn.userID = userIDHeader

	clientIDHeader := req.Header.Get("clientId")
	clientID, err := strconv.ParseInt(clientIDHeader, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	rtn.clientID = clientID
	if req.Header.Get("hashed") == "true" {
		rtn.hashed = true
	} else {
		rtn.hashed = false
	}
	//fmt.Println(clientIDHeader)
	//fmt.Println(userIDHeader)
	return rtn
}

func getAuth(req *http.Request) *uoauth.Oauth {
	changeHeader := getHeaders(req)
	auth := new(uoauth.Oauth)
	auth.Token = changeHeader.token
	auth.ClientID = changeHeader.clientID
	auth.UserID = changeHeader.userID
	auth.Hashed = changeHeader.hashed
	if os.Getenv("OAUTH2_VALIDATION_URI") != "" {
		auth.ValidationURL = os.Getenv("OAUTH2_VALIDATION_URI")
	} else {
		auth.ValidationURL = "http://localhost:3000/rs/token/validate"
	}
	return auth
}
