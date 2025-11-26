package growel

import (
	"net/http"
)

type Context struct {
	W      http.ResponseWriter
	R      http.Request
	Params map[string]string
}

type Handler func(c *Context)

func (h Handler) ServeHTTP(w http.ResponseWriter, r http.Request) {
	h(&Context{
		W:      w,
		R:      r,
		Params: make(map[string]string),
	})
}
