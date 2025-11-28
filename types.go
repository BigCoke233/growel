package growel

import "net/http"

type Context struct {
	W      http.ResponseWriter
	R      *http.Request
	Params map[string]string
	Querys map[string]string
}

type Handler func(c *Context)

type Route struct {
	method  string
	parts   []string
	handler Handler
}

type Router struct {
	routes []Route
}

type Engine struct {
	router *Router
}
