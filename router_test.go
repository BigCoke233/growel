package growel

import (
	"testing"
)

func dummy(c *Context) {}

func TestRouterStaticMatch(t *testing.T) {
	r := NewRouter()
	r.Add("GET", "/hello", dummy)

	h, params, _ := r.Find("GET", "/hello")
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

	h, _, _ := r.Find("GET", "/nothello")
	if h != nil {
		t.Fatal("expected nil handler, got non-nil")
	}
}

func TestRouterParamMatch(t *testing.T) {
	r := NewRouter()
	r.Add("GET", "/user/:id", dummy)

	h, params, _ := r.Find("GET", "/user/42")
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

	h, _, _ := r.Find("GET", "/user")
	if h != nil {
		t.Fatal("expected nil handler, got non-nil")
	}
}

func TestRouterQueryMatch(t *testing.T) {
	r := NewRouter()
	r.Add("GET", "/user", dummy)

	h, _, query := r.Find("GET", "/user?id=42")
	if h == nil {
		t.Fatal("expected handler, got nil")
	}
	if query["id"] != "42" {
		t.Fatalf("expected id=42, got %v", query)
	}
}
