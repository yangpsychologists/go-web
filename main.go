package main


func main() {

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
