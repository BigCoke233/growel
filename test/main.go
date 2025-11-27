package main

import (
	"github.com/bigcoke233/growel/v2"
	"strconv"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Group string `json:"group"`
}

var Users = []User{
	{ID: 1, Name: "Alice Chen", Group: "admin"},
	{ID: 2, Name: "Bob Wu", Group: "admin"},

	{ID: 3, Name: "Cynthia Lin", Group: "staff"},
	{ID: 4, Name: "Daniel Zhao", Group: "staff"},
	{ID: 5, Name: "Emma Huang", Group: "staff"},

	{ID: 6, Name: "Felix Sun", Group: "editor"},
	{ID: 7, Name: "Grace Liu", Group: "editor"},

	{ID: 8, Name: "Henry Qiu", Group: "guest"},
	{ID: 9, Name: "Ivy Xu", Group: "guest"},
	{ID: 10, Name: "Jack Ma", Group: "guest"},

	{ID: 11, Name: "Karen He", Group: "beta"},
	{ID: 12, Name: "Leo Fang", Group: "beta"},

	{ID: 13, Name: "Mia Tang", Group: "test"},
	{ID: 14, Name: "Noah Jin", Group: "test"},
	{ID: 15, Name: "Olivia Ding", Group: "test"},

	{ID: 16, Name: "Peter Yan", Group: "ops"},
	{ID: 17, Name: "Queena Zhou", Group: "ops"},
	{ID: 18, Name: "Ryan Li", Group: "ops"},

	{ID: 19, Name: "Sophia Wang", Group: "support"},
	{ID: 20, Name: "Tony Du", Group: "support"},
}

func main() {
	api := growel.New()

	api.GET("/", func(c *growel.Context) {
		c.Plain(200, "This is home.")
	})

	api.GET("/hello", func(c *growel.Context) {
		c.JSON(200, map[string]string{
			"message": "Hello Growel!",
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

	api.Start(":8080")
}
