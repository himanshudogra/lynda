package main

import (
	"io"
	"net/http"
)

func upTown(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}

func youUp(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "catty catty catty")
}

func main() {

	http.HandleFunc("/", upTown)
	http.HandleFunc("/cat/", youUp)

	http.ListenAndServe(":9000", nil)
}

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
