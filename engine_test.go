package growel

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func setupEngine() *Engine {
	e := New()

	e.GET("/", func(c *Context) {
		c.Plain(200, "home")
	})

	e.GET("/hello", func(c *Context) {
		c.JSON(200, map[string]string{"msg": "hi"})
	})

	e.GET("/user/:id", func(c *Context) {
		c.JSON(200, map[string]string{"id": c.Params["id"]})
	})

	e.GET("/error", func(c *Context) {
		c.BadRequest("bad")
	})

	return e
}

func TestEngineStaticRoute(t *testing.T) {
	e := setupEngine()

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	if w.Body.String() != "home" {
		t.Fatalf("expected body 'home', got %q", w.Body.String())
	}
}

func TestEngineJSONRoute(t *testing.T) {
	e := setupEngine()

	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	if !strings.Contains(w.Body.String(), `"msg":"hi"`) {
		t.Fatalf("unexpected json: %s", w.Body.String())
	}
}

func TestEngineParamRoute(t *testing.T) {
	e := setupEngine()

	req := httptest.NewRequest("GET", "/user/99", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	if !strings.Contains(w.Body.String(), `"id":"99"`) {
		t.Fatalf("unexpected json: %s", w.Body.String())
	}
}

func TestEngineNotFound(t *testing.T) {
	e := setupEngine()

	req := httptest.NewRequest("GET", "/doesnotexist", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestEngineErrorResponse(t *testing.T) {
	e := setupEngine()

	req := httptest.NewRequest("GET", "/error", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Fatalf("expected 400, got %d", w.Code)
	}

	body := w.Body.String()
	if !strings.Contains(body, `"error":"Bad Request"`) {
		t.Fatalf("missing error text: %s", body)
	}

	if !strings.Contains(body, `"message":"bad"`) {
		t.Fatalf("missing message: %s", body)
	}

	if !strings.Contains(body, `"status":"400"`) {
		t.Fatalf("missing status code: %s", body)
	}
}
