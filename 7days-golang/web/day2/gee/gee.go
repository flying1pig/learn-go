package gee

import "net/http"

type HandlerFunc func(c *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) addRoute(method, pattern string, handle HandlerFunc) {
	key := method + "-" + pattern
	e.router.handlers[key] = handle
}

func (e *Engine) Get(pattern string, handle HandlerFunc) {
	e.addRoute("GET", pattern, handle)
}

func (e *Engine) POST(pattern string, handle HandlerFunc) {
	e.addRoute("POST", pattern, handle)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(w, r)
	e.router.handle(ctx)
}
