package growel

import (
	"encoding/json"
	"net/http"
	"strconv"
)

/* === Writer Functions */

// Plain writes plain text to ResponseWriter
// with code and string passed as parameters
// this automatically sets Content-Type to "text/plain; charset=utf-8"
func (c *Context) Plain(code int, data string) {
	c.W.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.W.WriteHeader(code)
	c.W.Write([]byte(data))
}

// JSON receives a map or struct and converts it to JSON
// then writes it to ResponseWriter with code passed as parameter
// this automatically sets Content-Type to "application/json; charset=utf-8"
func (c *Context) JSON(code int, data any) {
	c.W.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.W.WriteHeader(code)

	enc := json.NewEncoder(c.W)
	_ = enc.Encode(data)
}

// XML is unfinished and works almost the same as Plain()
// but it sets Content-Type to "application/xml; charset=utf-8"
func (c *Context) XML(code int, data string) {
	c.W.Header().Set("Content-Type", "application/xml; charset=utf-8")
	c.W.WriteHeader(code)
	c.W.Write([]byte(data))
}

// Error writes an error response to ResponseWriter
// with code and message passed as parameters
// this automatically sets Content-Type to "application/json; charset=utf-8"
//
// Response is a JSON string and contains 3 values:
// error, message, status
func (c *Context) Error(code int, msg string) {
	c.JSON(code, map[string]string{
		"error":   http.StatusText(code),
		"message": msg,
		"status":  strconv.Itoa(code),
	})
}

// NotFound is a shortcut to write 404 Error to response writer
// this is based on Error function
func (c *Context) NotFound(msg string) {
	c.Error(http.StatusNotFound, msg)
}

// BadRequest is a shortcut to write 400 Error to response writer
// this is based on Error function
func (c *Context) BadRequest(msg string) {
	c.Error(http.StatusBadRequest, msg)
}

// Unauthorized is a shortcut to write 401 Error to response writer
// this is based on Error function
func (c *Context) Unauthorized(msg string) {
	c.Error(http.StatusUnauthorized, msg)
}

// Forbidden is a shortcut to write 403 Error to response writer
// this is based on Error function
func (c *Context) Forbidden(msg string) {
	c.Error(http.StatusForbidden, msg)
}

// InternalError is a shortcut to write 500 Error to response writer
// this is based on Error function
func (c *Context) InternalError(msg string) {
	c.Error(http.StatusInternalServerError, msg)
}

/* === Cookies === */

func (c *Context) Cookies() []*http.Cookie {
	return c.R.Cookies()
}

func (c *Context) Cookie(name string) *http.Cookie {
	for _, cookie := range c.R.Cookies() {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}

func (c *Context) SetCompleteCookie(cookie *http.Cookie) {
	http.SetCookie(c.W, cookie)
}

func (c *Context) SetCookie(name string, value string) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	c.SetCompleteCookie(&cookie)
}

/* === Request Data === */

// Context.FormValue gets value from request form
// form contains URL querys and POST Form value
//
// alias of 'c.Form[key]`
func (c *Context) FormValue(key string) string {
	return c.Form.Get(key)
}

// Context.PostFormValue works similarly to Context.FormValue
// but it does NOT contain URL querys, only POST form
func (c *Context) PostFormValue(key string) string {
	return c.R.PostFormValue(key)
}

// Context.Query gets value from request URL query
//
// alias of 'c.Querys[key]'
func (c *Context) Query(key string) string {
	return c.Querys.Get(key)
}

// Context.BindJSON takes the container for JSON data
// decodes request body and then writes data to container
func (c *Context) BindJSON(v any) error {
	c.R.Body = http.MaxBytesReader(c.W, c.R.Body, 2<<20)
	dec := json.NewDecoder(c.R.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(v)
}

/* === Parameter Type Helpers === */

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) ParamInt(key string) int {
	val, err := strconv.Atoi(c.Params[key])
	if err != nil {
		L.Error(err, "Context.ParamInt - cannot convert string to integer")
		return -1
	}
	return val
}

func (c *Context) ParamFloat(key string) float64 {
	val, err := strconv.ParseFloat(c.Params[key], 64)
	if err != nil {
		L.Error(err, "Context.ParamFloat - cannot convert string to float")
		return -1
	}
	return val
}

func (c *Context) ParamBool(key string) bool {
	val, err := strconv.ParseBool(c.Params[key])
	if err != nil {
		L.Error(err, "Context.ParamBool - cannot convert string to bool")
		return false
	}
	return val
}
