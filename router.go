package growel

import (
	"strings"
)

type route struct {
	method  string
	parts   []string
	handler Handler
}

type Router struct {
	routes []route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Add(method, path string, h Handler) {
	parts := strings.Split(path, "/")
	r.routes = append(r.routes, route{method, parts, h})
}

func (r *Router) Find(method, path string) (Handler, map[string]string) {
	parts := strings.Split(path, "/")

	for _, rt := range r.routes {
		// match HTTP method
		if rt.method != method {
			continue
		}

		// match path part number
		if len(rt.parts) != len(parts) {
			continue
		}

		// compare each segment
		params := make(map[string]string)
		matched := true
		for i, seg := range rt.parts {
			// parse dynamic parameter
			if strings.HasPrefix(seg, ":") {
				params[seg[1:]] = parts[i]
			}

			// static mismatch
			if seg != parts[i] {
				matched = false
				break
			}
		}
		if matched {
			return rt.handler, params
		}
	}
	return nil, nil
}
