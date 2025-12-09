package growel

import (
	"net/http"
)

// Context.Cookies returns all cookies from the request.
func (c *Context) Cookies() []*http.Cookie {
	return c.R.Cookies()
}

// Context.Cookie returns the cookie with the given name from the request.
func (c *Context) Cookie(name string) *http.Cookie {
	for _, cookie := range c.R.Cookies() {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}

// Context.SetCompleteCookie is the alias for `http.SetCookie(reponseWriter, ...)`
func (c *Context) SetCompleteCookie(cookie *http.Cookie) {
	http.SetCookie(c.W, cookie)
}

// Context.SetCookie is a shortcut to create a simple cookie,
// only name and value required.
func (c *Context) SetCookie(name string, value string) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	c.SetCompleteCookie(&cookie)
}
