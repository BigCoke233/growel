package growel

import (
	"encoding/json"
	"net/http"
)

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
