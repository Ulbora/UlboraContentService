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
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello Wolrd")
	fmt.Println("It works!!!")
	router := mux.NewRouter()
	router.HandleFunc("/rs/content", handleContent).Methods("POST", "PUT")
	router.HandleFunc("/rs/content/{id}/{clientId}", handleContentGet).Methods("GET", "DELETE")
	router.HandleFunc("/rs/contentList/{clientId}", handleContentList).Methods("GET")
	http.ListenAndServe(":3008", router)
}

func handleContent(res http.ResponseWriter, req *http.Request) {

}

func handleContentGet(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Print("id is: ")
	fmt.Print(id)
	var rtn = []byte("success")

	res.Write(rtn)
}

func handleContentList(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Print("id is: ")
	fmt.Print(id)
	var rtn = []byte("success")

	res.Write(rtn)
}
