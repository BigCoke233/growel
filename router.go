package growel

import (
	"strings"
)

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) splitPath(path string) []string {
	return strings.Split(strings.Trim(path, "/"), "/")
}

func (r *Router) Add(method string, path string, h Handler) {
	parts := r.splitPath(path)
	r.routes = append(r.routes, Route{method, parts, h})
}

func (r *Router) Find(method string, path string) (Handler, map[string]string) {
	parts := r.splitPath(path)

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
				continue
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
