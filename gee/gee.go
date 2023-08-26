package gee

import (
	"log"
	"net/http"
)

// HandlerFunc 类型，默认处理 HTTP 请求的函数类型，接收一个 *Context 对象作为参数。
type HandlerFunc func(*Context)

// Engine 结构体，定义整个框架的核心部分。Engine 结构体包含一个 router 字段，指向 router 结构体的指针。
type Engine struct {
	router *router
}

// New 函数，用于创建一个新的 Engine 对象。New() 函数返回一个指向 Engine 结构体的指针，并初始化了 router 字段。
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute 方法，用于向 router 添加路由规则。它接收 HTTP 方法、路径模式和处理函数作为参数，并将路由规则添加到 router 中。
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router.addRouter(method, pattern, handler)
}

// GET 方法定义，用于添加 GET 请求的路由规则。这些方法内部调用了 addRoute() 方法。
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 方法定义，用于添加 GET 和 POST 请求的路由规则。这些方法内部调用了 addRoute() 方法。
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// ServeHTTP() 方法，该方法是 Engine 类型实现了 http.Handler 接口的方法。
// 它接收 http.ResponseWriter 对象和 *http.Request 对象作为参数，并创建一个新的 Context 对象，
// 然后调用 router 的 handle() 方法来处理请求。
func (engine *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	c := newContext(resp, req)
	engine.router.handle(c)
}

// Run 方法定义，用于启动 HTTP 服务器。它接收一个地址字符串作为参数，调用 http.ListenAndServe() 函数来监听指定地址，
// 并将 engine 作为参数传递给 http.ListenAndServe() 函数。
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
