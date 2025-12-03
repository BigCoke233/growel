package growel

import (
	"net/http"
	"net/url"
)

type Context struct {
	W      http.ResponseWriter
	R      *http.Request
	Params map[string]string
	Querys url.Values
	Form   url.Values
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

type Group struct {
	prefix string
	engine *Engine
}
