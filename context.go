package growel

import (
	"net/http"
)

type Context struct {
	W      http.ResponseWriter
	R      *http.Request
	Params map[string]string
}

type Handler func(c *Context)
