package growel

import (
	"encoding/json"
	"net/http"
	"strconv"
)

/* Form: both URL Query and POST Form */

// FormValue gets value from request form
// form contains both URL querys and POST Form value
//
// alias of 'c.Form[key]`
func (c *Context) FormValue(key string) string {
	return c.Form.Get(key)
}

// FormInt gets integer value from request form
// form contains both URL querys and POST Form value
func (c *Context) FormInt(key string) int {
	val, err := strconv.Atoi(c.FormValue(key))
	if err != nil {
		L.Error(err, "Context.FormInt - cannot convert string to integer")
		return -1
	}
	return val
}

// FormFloat64 gets float64 value from request form
// form contains both URL querys and POST Form value
func (c *Context) FormFloat(key string) float64 {
	val, err := strconv.ParseFloat(c.FormValue(key), 64)
	if err != nil {
		L.Error(err, "Context.FormFloat64 - cannot convert string to float64")
		return -1
	}
	return val
}

// FormBool gets boolean value from request form
// form contains both URL querys and POST Form value
func (c *Context) FormBool(key string) bool {
	val, err := strconv.ParseBool(c.FormValue(key))
	if err != nil {
		L.Error(err, "Context.FormBool - cannot convert string to boolean")
		return false
	}
	return val
}

/* PostForm: Only POST form data */

// PostFormValue works similarly to Context.FormValue
// but it does NOT contain URL querys, only POST form
func (c *Context) PostFormValue(key string) string {
	return c.R.PostFormValue(key)
}

// PostFormInt gets integer value from request form
// form contains both URL querys and POST Form value
func (c *Context) PostFormInt(key string) int {
	val, err := strconv.Atoi(c.PostFormValue(key))
	if err != nil {
		L.Error(err, "Context.PostFormInt - cannot convert string to integer")
		return -1
	}
	return val
}

// PostFormFloat gets float64 value from request form
// form contains both URL querys and POST Form value
func (c *Context) PostFormFloat(key string) float64 {
	val, err := strconv.ParseFloat(c.PostFormValue(key), 64)
	if err != nil {
		L.Error(err, "Context.PostFormFloat - cannot convert string to float64")
		return -1
	}
	return val
}

// PostFormBool gets boolean value from request form
// form contains both URL querys and POST Form value
func (c *Context) PostFormBool(key string) bool {
	val, err := strconv.ParseBool(c.PostFormValue(key))
	if err != nil {
		L.Error(err, "Context.PostFormBool - cannot convert string to boolean")
		return false
	}
	return val
}

/* Query: only URL query data */

// Query gets value from request URL query
//
// alias for 'c.Querys[key]'
func (c *Context) Query(key string) string {
	return c.Querys.Get(key)
}

// QueryInt gets int value from request URL query
func (c *Context) QueryInt(key string) int {
	val, err := strconv.Atoi(c.Query(key))
	if err != nil {
		L.Error(err, "Context.QueryInt - cannot convert string to int")
		return -1
	}
	return val
}

// QueryFloat gets float64 value from request URL query
func (c *Context) QueryFloat(key string) float64 {
	val, err := strconv.ParseFloat(c.Query(key), 64)
	if err != nil {
		L.Error(err, "Context.QueryFloat - cannot convert string to float64")
		return -1
	}
	return val
}

// QueryBool gets boolean value from request URL query
func (c *Context) QueryBool(key string) bool {
	val, err := strconv.ParseBool(c.Query(key))
	if err != nil {
		L.Error(err, "Context.QueryBool - cannot convert string to boolean")
		return false
	}
	return val
}

/* BindJSON: data from request body */

// Context.BindJSON takes the container for JSON data
// decodes request body and then writes data to container
func (c *Context) BindJSON(v any) error {
	c.R.Body = http.MaxBytesReader(c.W, c.R.Body, 2<<20)
	dec := json.NewDecoder(c.R.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(v)
}
