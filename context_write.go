package growel

import (
	"encoding/json"
)

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
