package growel

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGroupCallback(t *testing.T) {
	e := New()

	called := false

	e.Group("/user", func(g *Group) {
		if g.prefix != "/user" {
			t.Fatalf("expected prefix '/user', got '%s'", g.prefix)
		}
		if g.engine != e {
			t.Fatal("group.engine does not point to Engine")
		}
		called = true
	})

	if !called {
		t.Fatal("Group callback not invoked")
	}
}

func TestGroupRouteRegistration(t *testing.T) {
	e := New()

	e.Group("/api", func(api *Group) {
		api.GET("/ping", func(c *Context) {})
		api.POST("/user", func(c *Context) {})
		api.PUT("/v1/item/", func(c *Context) {})
		api.DELETE("/v1/item/42", func(c *Context) {})
	})

	routes := e.router.routes

	if len(routes) != 4 {
		t.Fatalf("expected 4 routes, got %d", len(routes))
	}

	cases := []struct {
		method string
		path   string
	}{
		{"GET", "/api/ping"},
		{"POST", "/api/user"},
		{"PUT", "/api/v1/item"},
		{"DELETE", "/api/v1/item/42"},
	}

	for _, want := range cases {
		found := false

		for _, r := range routes {
			joined := "/" + joinParts(r.parts)
			if r.method == want.method && joined == want.path {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("missing expected route: %s %s", want.method, want.path)
		}
	}
}

func TestGroupServeHTTP(t *testing.T) {
	e := New()

	e.Group("/hello", func(h *Group) {
		h.GET("/world", func(c *Context) {
			c.Plain(200, "ok")
		})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello/world", nil)

	e.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	if body := w.Body.String(); body != "ok\n" && body != "ok" {
		t.Fatalf("expected body 'ok', got '%s'", body)
	}
}

// helper: convert []string{"api","ping"} → "api/ping"
func joinParts(parts []string) string {
	out := ""
	for i, p := range parts {
		if i > 0 {
			out += "/"
		}
		out += p
	}
	return out
}
