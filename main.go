package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
func indexHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "URL.Path = %q", req.URL.Path)
}

func helloHandler(resp http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(resp, "Header[%q] = %q\n", k, v)
	}
}

/*
1. GET http://127.0.0.1:8080/
	URL.Path = "/"

2. GET http://127.0.0.1:8080/hello
	Header["Accept"] = ["* /*"]
	Header["Postman-Token"] = ["34ed4dd7-4ba8-48c5-aed2-748504ceb54a"]
	Header["Accept-Encoding"] = ["gzip, deflate, br"]
	Header["Connection"] = ["keep-alive"]
*/
