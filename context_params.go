package growel

import (
	"strconv"
)

// Context.Param() is the alias for Context.Params[key]
//
// In growel, "params" refer to dynamic parameters in the URL path.
// as in `/user/:id`, in which case the value of `id` is stored in `Context.Params["id"]`.
func (c *Context) Param(key string) string {
	return c.Params[key]
}

// Context.ParamInt() converts Context.Param(key) to int and returns it
//
// In growel, "params" refer to dynamic parameters in the URL path.
// as in `/user/:id`, in which case the value of `id` is stored in `Context.Params["id"]`.
func (c *Context) ParamInt(key string) int {
	val, err := strconv.Atoi(c.Params[key])
	if err != nil {
		L.Error(err, "Context.ParamInt - cannot convert string to integer")
		return -1
	}
	return val
}

// Context.ParamFloat() converts Context.Param(key) to float64 and returns it
//
// In growel, "params" refer to dynamic parameters in the URL path.
// as in `/user/:id`, in which case the value of `id` is stored in `Context.Params["id"]`.
func (c *Context) ParamFloat(key string) float64 {
	val, err := strconv.ParseFloat(c.Params[key], 64)
	if err != nil {
		L.Error(err, "Context.ParamFloat - cannot convert string to float")
		return -1
	}
	return val
}

// Context.ParamBool() converts Context.Param(key) to bool and returns it
//
// In growel, "params" refer to dynamic parameters in the URL path.
// as in `/user/:id`, in which case the value of `id` is stored in `Context.Params["id"]`.
func (c *Context) ParamBool(key string) bool {
	val, err := strconv.ParseBool(c.Params[key])
	if err != nil {
		L.Error(err, "Context.ParamBool - cannot convert string to bool")
		return false
	}
	return val
}
