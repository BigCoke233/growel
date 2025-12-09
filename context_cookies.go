package growel

import (
	"net/http"
	"time"
)

// Cookies returns all cookies from the request.
func (c *Context) Cookies() []*http.Cookie {
	return c.R.Cookies()
}

// Cookie returns the cookie with the given name from the request.
func (c *Context) Cookie(name string) *http.Cookie {
	for _, cookie := range c.R.Cookies() {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}

// HasCookie returns true if the request contains a cookie with the given
// name, returns false if not.
func (c *Context) HasCookie(name string) bool {
	for _, cookie := range c.R.Cookies() {
		if cookie.Name == name {
			return true
		}
	}
	return false
}

// SetCompleteCookie is the alias for `http.SetCookie(reponseWriter, ...)`
func (c *Context) SetCompleteCookie(cookie *http.Cookie) {
	http.SetCookie(c.W, cookie)
}

// SetCookie is a shortcut to create a cookie with only name and value required.
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

// DeleteCookie sets cookie value to empty string and max age to -1,
// practically deleting it.
func (c *Context) DeleteCookie(name string) {
	cookie := http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	c.SetCompleteCookie(&cookie)
}

// ClearCookies deletes all cookies from the request.
func (c *Context) ClearCookies() {
	for _, cookie := range c.R.Cookies() {
		c.DeleteCookie(cookie.Name)
	}
}

/* BELOW ARE HELPER FUNCTIONS TO MAKE COOKIE SETTING LESS VERBOSE  */

// CookiePath sets path for cookie with the given name.
func (c *Context) CookiePath(name string, path string) {
	for _, cookie := range c.R.Cookies() {
		if cookie.Name == name {
			cookie.Path = path
			c.SetCompleteCookie(cookie)
			return
		}
	}
}

// CookieAge sets max age for cookie with the given name.
func (c *Context) CookieMaxAge(name string, age int) {
	for _, cookie := range c.R.Cookies() {
		if cookie.Name == name {
			cookie.MaxAge = age
			c.SetCompleteCookie(cookie)
			return
		}
	}
}

// CookieSecure sets secure flag for cookie with the given name.
func (c *Context) CookieSecure(name string, secure bool) {
	for _, cookie := range c.R.Cookies() {
		if cookie.Name == name {
			cookie.Secure = secure
			c.SetCompleteCookie(cookie)
			return
		}
	}
}

// CookieSameSite sets SameSite flag for cookie with the given name.
func (c *Context) CookieSameSite(name string, sameSite http.SameSite) {
	for _, cookie := range c.R.Cookies() {
		if cookie.Name == name {
			cookie.SameSite = sameSite
			c.SetCompleteCookie(cookie)
			return
		}
	}
}

// CookieHttpOnly sets HttpOnly flag for cookie with the given name.
func (c *Context) CookieHttpOnly(name string, httpOnly bool) {
	for _, cookie := range c.R.Cookies() {
		if cookie.Name == name {
			cookie.HttpOnly = httpOnly
			c.SetCompleteCookie(cookie)
			return
		}
	}
}

// CookieDomain sets domain for cookie with the given name.
func (c *Context) CookieDomain(name string, domain string) {
	for _, cookie := range c.R.Cookies() {
		if cookie.Name == name {
			cookie.Domain = domain
			c.SetCompleteCookie(cookie)
			return
		}
	}
}

// CookieExpires sets expiration time for cookie with the given name.
func (c *Context) CookieExpires(name string, expires time.Time) {
	for _, cookie := range c.R.Cookies() {
		if cookie.Name == name {
			cookie.Expires = expires
			c.SetCompleteCookie(cookie)
			return
		}
	}
}
