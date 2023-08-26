package gee

import "net/http"

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRouter(method string, pattern string, hadler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = hadler
}
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if hadler, ok := r.handlers[key]; ok {
		hadler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
