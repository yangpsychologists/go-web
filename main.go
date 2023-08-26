package main

import (
	"go-web/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	// /hello?name=geetutu
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":8080")
}

/*
1. GET http://127.0.0.1:8080/
	<h1>Hello Gee</h1>

2. GET http://127.0.0.1:8080/hello?name=geektutu
	hello geektutu, you're at /hello

3. POST http://127.0.0.1:8080/login?username=geektutu&password=1234
	{
    "password": "1234",
    "username": "geektutu"
	}

4. GET http://127.0.0.1:8080/xxx
	404 NOT FOUND: /xxx

*/
