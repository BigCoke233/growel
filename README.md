# Growel

![](https://github.com/BigCoke233/growel/actions/workflows/go.yml/badge.svg)

**Minimal, headless, zero-dependency Web framework written in Go.** This is toy software I develop to learn, so do not use in production.

## Usage

```go
e := growel.New()

e.GET("/hello", func(c *growel.Context) {
	c.JSON(200, map[string]string{
		"message": "Hello, Growel!"
	})
})

e.Group("/api", func(api *growel.Group) {
	api.Group("/user", func(user *growel.Group) {
		user.GET("/", func(c *growel.Context) {
			c.JSON(200, Users)
		})

		user.GET("/:uid", func(c *growel.Context) {
			uid, err := strconv.Atoi(c.Params["uid"])
			if err != nil {
				c.BadRequest("Invalid user ID")
				return
			}

			for _, user := range Users {
				if user.ID == uid {
					c.JSON(200, user)
					return
				}
			}

			c.NotFound("User not found")
		})
	})
})
```

## Roadmap

- [ ] Authorization with token
- [ ] Panic recovery
- [ ] CORS support
- [ ] Custom 404 handler
- [ ] `Context.Redirect()`
- [ ] File Upload
- [ ] Serve static files
- [ ] `Engine.shutdown()`
- [ ] TLS
- [x] Wildcard routing `*`
- [x] Parameter type support, `ParamInt()` `ParamString()`
