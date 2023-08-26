package gee

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc 类型，定义默认处理 HTTP 请求的函数， 接收一个 http.ResponseWriter 对象和一个 *http.Request 对象作为参数。
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 结构体，整个框架的核心部分。Engine 结构体包含一个 router 字段，类型为 map；用于存储不同 HTTP 方法和路径模式对应的处理函数。
type Engine struct {
	router map[string]HandlerFunc
}

// New 函数，用于创建一个新的 Engine 对象。New() 函数返回一个指向 Engine 结构体的指针。
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute 方法定义，用于向 router 添加路由规则。它接收 HTTP 方法、路径模式和处理函数作为参数，将路由规则添加到 router 映射中。
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	engine.router[key] = handler
}

// GET 方法定义，用于添加 GET 请求的路由规则。这些方法内部调用了 addRoute() 方法。
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 方法定义，用于添加 GET 和 POST 请求的路由规则。这些方法内部调用了 addRoute() 方法。
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// ServeHTTP 方法定义，该方法是 Engine 类型实现了 http.Handler 接口的方法。
// 它接收 http.ResponseWriter 对象和 *http.Request 对象作为参数，根据请求的方法和路径查找对应的处理函数，并执行该处理函数。
func (engine *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(resp, req)
	} else {
		fmt.Fprintf(resp, "404 NOT FOUND: %s\n", req.URL)
	}
}

// Run 方法定义，用于启动 HTTP 服务器。它接收一个地址字符串作为参数，调用 http.ListenAndServe() 函数来监听指定地址，
// 并将 engine 作为参数传递给 http.ListenAndServe() 函数。
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
