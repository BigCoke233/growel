package main

import (
	"fmt"

	"github.com/bigcoke233/growel/v2"
)

func main() {
	api := growel.New()

	api.GET("/hello", func(c *growel.Context) {
		c.W.Write([]byte("Hello, World!"))
	})

	api.GET("/", func(c *growel.Context) {
		c.W.Write([]byte("This is home."))
	})

	api.GET("/test", func(c *growel.Context) {
		c.W.Write([]byte("testing"))
	})

	api.GET("/test/hello", func(c *growel.Context) {
		c.W.Write([]byte("testing hello"))
	})

	api.GET("/test/:uid", func(c *growel.Context) {
		c.W.Write([]byte(fmt.Sprintf("Testing from %s", c.Params["uid"])))
	})

	api.GET("/test/:uid/info/:sid", func(c *growel.Context) {
		c.W.Write([]byte(fmt.Sprintf("Testing from %s info %s", c.Params["uid"], c.Params["sid"])))
	})

	api.Start(":8080")
}
