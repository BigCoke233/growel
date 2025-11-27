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

func TestRouterSplitPath(t *testing.T) {
	r := NewRouter()

	tests := []struct {
		input string
		want  []string
	}{
		{"/", []string{""}},
		{"/hello", []string{"hello"}},
		{"/hello/world", []string{"hello", "world"}},
		{"//hello//world//", []string{"hello", "", "world"}},
	}

	for _, tt := range tests {
		got := r.splitPath(tt.input)
		if len(got) != len(tt.want) {
			t.Fatalf("splitPath(%q) length mismatch: got %v want %v", tt.input, got, tt.want)
		}
		for i := range got {
			if got[i] != tt.want[i] {
				t.Fatalf("splitPath(%q) mismatch: got %v want %v", tt.input, got, tt.want)
			}
		}
	}
}
