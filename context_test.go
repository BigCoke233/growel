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

func TestContextCookies(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	r.AddCookie(&http.Cookie{Name: "a", Value: "1"})
	r.AddCookie(&http.Cookie{Name: "b", Value: "2"})

	c := &Context{R: r}

	cookies := c.Cookies()
	if len(cookies) != 2 {
		t.Fatalf("expected 2 cookies, got %d", len(cookies))
	}
}

func TestContextCookieFound(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	r.AddCookie(&http.Cookie{Name: "token", Value: "xyz"})

	c := &Context{R: r}

	ck := c.Cookie("token")
	if ck == nil {
		t.Fatalf("expected cookie token, got nil")
	}
	if ck.Value != "xyz" {
		t.Fatalf("expected value xyz, got %s", ck.Value)
	}
}

func TestContextCookieNotFound(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	r.AddCookie(&http.Cookie{Name: "session", Value: "abc"})

	c := &Context{R: r}

	ck := c.Cookie("missing")
	if ck != nil {
		t.Fatalf("expected nil, got %#v", ck)
	}
}

func TestContextSetCompleteCookie(t *testing.T) {
	w := httptest.NewRecorder()
	c := &Context{W: w}

	cookie := &http.Cookie{
		Name:  "x",
		Value: "y",
		Path:  "/",
	}

	c.SetCompleteCookie(cookie)

	res := w.Result()
	cs := res.Cookies()

	if len(cs) != 1 {
		t.Fatalf("expected 1 cookie, got %d", len(cs))
	}

	ck := cs[0]
	if ck.Name != "x" || ck.Value != "y" {
		t.Fatalf("unexpected cookie: %#v", ck)
	}
}

func TestContextSetCookie(t *testing.T) {
	w := httptest.NewRecorder()
	c := &Context{W: w}

	c.SetCookie("uid", "100")

	res := w.Result()
	cs := res.Cookies()

	if len(cs) != 1 {
		t.Fatalf("expected 1 cookie, got %d", len(cs))
	}

	ck := cs[0]

	if ck.Name != "uid" || ck.Value != "100" {
		t.Fatalf("unexpected cookie: %#v", ck)
	}
	if ck.Path != "/" {
		t.Fatalf("wrong Path: %s", ck.Path)
	}
	if ck.MaxAge != 3600 {
		t.Fatalf("wrong MaxAge: %d", ck.MaxAge)
	}
	if !ck.HttpOnly {
		t.Fatalf("HttpOnly should be true")
	}
	if !ck.Secure {
		t.Fatalf("Secure should be true")
	}
	if ck.SameSite != http.SameSiteLaxMode {
		t.Fatalf("wrong SameSite: %v", ck.SameSite)
	}
}
