package growel

import (
	"slices"
	"strings"
)

func NewRouter() *Router {
	return &Router{}
}

// Router.Add adds a route with specified method, path and handler
// it splits path into segments for matching
func (r *Router) Add(method string, path string, h Handler) {
	parts := splitPath(path)
	r.routes = append(r.routes, Route{method, parts, h})
}

// Router.Find finds route with same HTTP method and path
// returns the handler and dynamic parameters
//
// handler is a function that wraps Context
// params is a map of parameters from dynamic routing
func (r *Router) Find(method string, path string) (
	handler Handler, params map[string]string) {
	parts := splitPath(path)

	for _, rt := range r.routes {
		// match HTTP method
		if rt.method != method {
			continue
		}

		// match path part number
		// if has ** wildcard, ignore part number
		hasWildcard := slices.Contains(rt.parts, "**")
		if !hasWildcard && len(rt.parts) != len(parts) {
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

			// wildcard
			if seg == "*" {
				continue
			}
			if seg == "**" {
				break
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

/* === helpers ===  */

// splitPath splits path into segments for matching
func splitPath(path string) []string {
	return strings.Split(strings.Trim(path, "/"), "/")
}
