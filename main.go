package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func indexHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "URL.Path = %q\n", req.URL.Path)
}
func helloHandler(resp http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(resp, "Header[%q] = %q\n", k, v)
	}
}

/*
1. GET 127.0.0.1:8080/
	URL.Path = "/"

2. GET 127.0.0.1:8080/hello
	Header["Accept"] = ["* /*"]
	Header["Postman-Token"] = ["c2709648-34a9-4bc3-b012-1fb8b97e55bf"]
	Header["Accept-Encoding"] = ["gzip, deflate, br"]
	Header["Connection"] = ["keep-alive"]
	Header["User-Agent"] = ["PostmanRuntime/7.32.3"]
*/
