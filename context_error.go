package growel

import (
	"net/http"
	"strconv"
)

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
