package growel

import (
	"testing"
)

func dummy(c *Context) {}

func TestRouterStaticMatch(t *testing.T) {
	r := NewRouter()
	r.Add("GET", "/hello", dummy)

	h, params := r.Find("GET", "/hello")
	if h == nil {
		t.Fatal("expected handler, got nil")
	}
	if len(params) != 0 {
		t.Fatalf("expected no params, got %v", params)
	}
}

func TestRouterStaticNoMatch(t *testing.T) {
	r := NewRouter()
	r.Add("GET", "/hello", dummy)

	h, _ := r.Find("GET", "/nothello")
	if h != nil {
		t.Fatal("expected nil handler, got non-nil")
	}
}

func TestRouterParamMatch(t *testing.T) {
	r := NewRouter()
	r.Add("GET", "/user/:id", dummy)

	h, params := r.Find("GET", "/user/42")
	if h == nil {
		t.Fatal("expected handler, got nil")
	}
	if params["id"] != "42" {
		t.Fatalf("expected id=42, got %v", params)
	}
}

func TestRouterParamMismatch(t *testing.T) {
	r := NewRouter()
	r.Add("GET", "/user/:id", dummy)

	h, _ := r.Find("GET", "/user")
	if h != nil {
		t.Fatal("expected nil handler, got non-nil")
	}
}

func TestRouterWildcardMatch(t *testing.T) {
	r := NewRouter()
	r.Add("GET", "/static/*", dummy)

	h, params := r.Find("GET", "/static/app.css")
	if h == nil {
		t.Fatal("expected handler, got nil")
	}
	if len(params) != 0 {
		t.Fatalf("expected no params, got %v", params)
	}
}

func TestRouterDoubleWildcardMatch(t *testing.T) {
	r := NewRouter()
	r.Add("GET", "/static/**", dummy)

	h, params := r.Find("GET", "/static/css/app.css")
	if h == nil {
		t.Fatal("expected handler, got nil")
	}
	if len(params) != 0 {
		t.Fatalf("expected no params, got %v", params)
	}
}

func TestRouterWildcardNoMatch(t *testing.T) {
	r := NewRouter()
	r.Add("GET", "/static/*", dummy)

	h, _ := r.Find("GET", "/assets/app.css")
	if h != nil {
		t.Fatal("expected nil handler, got non-nil")
	}
}

func TestRouterGlobalWildcardMatch(t *testing.T) {
	r := NewRouter()
	r.Add("GET", "**", dummy)

	h, params := r.Find("GET", "/anything/really")
	if h == nil {
		t.Fatal("expected handler, got nil")
	}
	if len(params) != 0 {
		t.Fatalf("expected no params, got %v", params)
	}
}

func TestRouterGlobalWildcardWithOtherRoutes(t *testing.T) {
	r := NewRouter()
	r.Add("GET", "/hello", dummy)
	r.Add("GET", "**", dummy)

	// Exact route should win
	h, _ := r.Find("GET", "/hello")
	if h == nil {
		t.Fatal("expected exact handler, got nil")
	}

	// Fallback wildcard should catch everything else
	h2, params := r.Find("GET", "/not/hello")
	if h2 == nil {
		t.Fatal("expected fallback wildcard handler, got nil")
	}
	if len(params) != 0 {
		t.Fatalf("expected no params, got %v", params)
	}
}
