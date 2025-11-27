# Growel

**Minimal, headless, zero-dependency Web framework written in Go.** This is toy software I develop to learn, so do not use in production.

## Usage

```go
api := growel.New()

api.GET("/hello", func(c *growel.Context) {
	c.JSON(200, map[string]string{
		"message": "Hello, Growel!"
	})
})

api.GET("/user/", func(c *growel.Context) {
	c.JSON(200, Users)
})

api.GET("/user/:uid", func(c *growel.Context) {
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
```

## License

MIT
