package growel

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (c *Context) Plain(code int, data string) {
	c.W.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.W.WriteHeader(code)
	c.W.Write([]byte(data))
}

func (c *Context) JSON(code int, data any) {
	c.W.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.W.WriteHeader(code)

	enc := json.NewEncoder(c.W)
	_ = enc.Encode(data)
}

func (c *Context) Error(code int, msg string) {
	c.JSON(code, map[string]string{
		"error": http.StatusText(code),
		"message": msg,
		"status": strconv.Itoa(code),
	})
}

func (c *Context) NotFound(msg string) {
	c.Error(http.StatusNotFound, msg)
}

func (c *Context) BadRequest(msg string) {
	c.Error(http.StatusBadRequest, msg)
}

func (c *Context) Unauthorized(msg string) {
	c.Error(http.StatusUnauthorized, msg)
}

func (c *Context) InternalError(msg string) {
	c.Error(http.StatusInternalServerError, msg)
}
