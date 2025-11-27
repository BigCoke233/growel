package growel

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestContextPlain(t *testing.T) {
	w := httptest.NewRecorder()
	c := &Context{W: w}

	c.Plain(201, "ok")

	if w.Code != 201 {
		t.Fatalf("expected 201, got %d", w.Code)
	}
	if ct := w.Header().Get("Content-Type"); !strings.Contains(ct, "text/plain") {
		t.Fatalf("wrong content-type: %s", ct)
	}
	if w.Body.String() != "ok" {
		t.Fatalf("wrong body: %s", w.Body.String())
	}
}

func TestContextJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c := &Context{W: w}

	c.JSON(200, map[string]string{"x": "y"})

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if ct := w.Header().Get("Content-Type"); !strings.Contains(ct, "application/json") {
		t.Fatalf("unexpected content-type: %s", ct)
	}
	if !strings.Contains(w.Body.String(), `"x":"y"`) {
		t.Fatalf("wrong json: %s", w.Body.String())
	}
}

func TestContextError(t *testing.T) {
	w := httptest.NewRecorder()
	c := &Context{W: w}

	c.Error(http.StatusBadRequest, "bad input")

	if w.Code != 400 {
		t.Fatalf("expected 400, got %d", w.Code)
	}

	body := w.Body.String()
	if !strings.Contains(body, `"error":"Bad Request"`) {
		t.Fatalf("missing human readable status: %s", body)
	}
	if !strings.Contains(body, `"message":"bad input"`) {
		t.Fatalf("missing message: %s", body)
	}
	if !strings.Contains(body, `"status":"400"`) {
		t.Fatalf("missing status field: %s", body)
	}
}
