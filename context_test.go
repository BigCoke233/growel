package growel

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
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

func TestContextFormValue(t *testing.T) {

	// simulate query
	query, _ := url.ParseQuery("q=hello&page=2")

	// simulate POST form
	post := url.Values{}
	post.Set("username", "eltrac")
	post.Set("password", "123")

	req := httptest.NewRequest("POST", "/test?q=hello&page=2", nil)
	req.Form = make(url.Values)
	for k, v := range query {
		req.Form[k] = v
	}
	for k, v := range post {
		req.Form[k] = append(req.Form[k], v...)
	}

	c := &Context{
		R:      req,
		Form:   req.Form,
		Querys: req.URL.Query(),
	}

	if got := c.FormValue("q"); got != "hello" {
		t.Errorf("FormValue(q) = %s, want hello", got)
	}
	if got := c.FormValue("username"); got != "eltrac" {
		t.Errorf("FormValue(username) = %s, want eltrac", got)
	}
	if got := c.FormValue("password"); got != "123" {
		t.Errorf("FormValue(password) = %s, want 123", got)
	}
}

func TestContextPostFormValue(t *testing.T) {
	body := bytes.NewBufferString("x=1&y=2")
	req := httptest.NewRequest("POST", "/test?z=999", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.ParseForm()

	c := &Context{
		R:      req,
		Querys: req.URL.Query(),
		Form:   req.Form,
	}

	// PostFormValue should NOT include query
	if got := c.PostFormValue("z"); got != "" {
		t.Errorf("PostFormValue(z) should be empty, got %s", got)
	}

	if got := c.PostFormValue("x"); got != "1" {
		t.Errorf("PostFormValue(x) = %s, want 1", got)
	}
	if got := c.PostFormValue("y"); got != "2" {
		t.Errorf("PostFormValue(y) = %s, want 2", got)
	}
}

func TestContextQuery(t *testing.T) {
	req := httptest.NewRequest("GET", "/search?kw=hello&sort=asc", nil)

	c := &Context{
		R:      req,
		Querys: req.URL.Query(),
	}

	if got := c.Query("kw"); got != "hello" {
		t.Errorf("Query(kw) = %s, want hello", got)
	}

	if got := c.Query("sort"); got != "asc" {
		t.Errorf("Query(sort) = %s, want asc", got)
	}

	if got := c.Query("page"); got != "" {
		t.Errorf("Query(page) should be empty, got %s", got)
	}
}

func TestContextBindJSON(t *testing.T) {

	type User struct {
		Name  string `json:"name"`
		Level int    `json:"level"`
	}

	data := `{"name":"Eltrac","level":42}`
	req := httptest.NewRequest("POST", "/json", bytes.NewBufferString(data))

	w := httptest.NewRecorder()

	c := &Context{
		W: w,
		R: req,
	}

	var u User
	err := c.BindJSON(&u)

	if err != nil {
		t.Fatalf("BindJSON returned error: %v", err)
	}

	if u.Name != "Eltrac" {
		t.Errorf("BindJSON Name = %s, want Eltrac", u.Name)
	}
	if u.Level != 42 {
		t.Errorf("BindJSON Level = %d, want 42", u.Level)
	}
}

func TestBindJSONUnknownField(t *testing.T) {
	type User struct {
		Name string `json:"name"`
	}

	data := `{"name":"Eltrac", "extra":"BAD"}`
	req := httptest.NewRequest("POST", "/json", bytes.NewBufferString(data))
	w := httptest.NewRecorder()
	c := &Context{W: w, R: req}

	var u User
	err := c.BindJSON(&u)

	if err == nil {
		t.Fatalf("expected error for unknown field, got nil")
	}
}
