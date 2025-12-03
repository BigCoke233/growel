package growel

import (
	"strings"
)

/* constructor  */

// Engine.Group creates a Group
// pass the Group as a parameter to the callback function
// used to create route groups
func (e *Engine) Group(prefix string, fn func(*Group)) {
	g := &Group{
		prefix: prefix,
		engine: e,
	}

	fn(g)
}

/* Inherit Engine methods */

func (g *Group) GET(path string, handler Handler) {
	full := joinPath(g.prefix, path)
	g.engine.GET(full, handler)
}

func (g *Group) POST(path string, handler Handler) {
	full := joinPath(g.prefix, path)
	g.engine.POST(full, handler)
}

func (g *Group) DELETE(path string, handler Handler) {
	full := joinPath(g.prefix, path)
	g.engine.DELETE(full, handler)
}

func (g *Group) PUT(path string, handler Handler) {
	full := joinPath(g.prefix, path)
	g.engine.PUT(full, handler)
}

/* utils */

func joinPath(a, b string) string {
	if a == "/" {
		return a + strings.TrimPrefix(b, "/")
	}
	return strings.TrimSuffix(a, "/") + "/" + strings.TrimPrefix(b, "/")
}
