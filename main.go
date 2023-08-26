package main

import (
	"fmt"
	"go-web/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "URL.Path = %q\n", req.URL.Path)
	})
	r.GET("/hello", func(resp http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(resp, "Header[%q] = %q\n", k, v)
		}
	})
	r.Run(":8080")
}

/*
1. GET http://127.0.0.1:8080/
	URL.Path = "/"

2. GET http://127.0.0.1:8080/hello
	Header["Accept"] = ["* /*"]
	Header["Postman-Token"] = ["34ed4dd7-4ba8-48c5-aed2-748504ceb54a"]
	Header["Accept-Encoding"] = ["gzip, deflate, br"]
	Header["Connection"] = ["keep-alive"]

3. GET 127.0.0.1:8080/world
	404 NOT FOUND: /world

4. Log
	2023/08/27 02:54:54 Route  GET - /
	2023/08/27 02:54:54 Route  GET - /hello

*/
