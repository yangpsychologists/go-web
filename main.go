package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":8080", engine))
}

type Engine struct{}

func (engine *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(resp, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(resp, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(resp, "404 NOT FOUND: %s\n", req.URL)
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

3. GET http://127.0.0.1:8080/world
	404 NOT FOUND: /world
*/
