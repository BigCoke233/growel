package growel

import (
	"net/http"
)

type Api struct {
	router *Router
}

// create API engine

func New() *Api {
	var a Api
	a.router = NewRouter()
	return &a
}

// add routes

func (api *Api) GET(path string, h Handler) {
	api.router.Add("GET", path, h)
}

func (api *Api) POST(path string, h Handler) {
	api.router.Add("POST", path, h)
}

func (api *Api) PUT(path string, h Handler) {
	api.router.Add("PUT", path, h)
}

func (api *Api) DELETE(path string, h Handler) {
	api.router.Add("DELETE", path, h)
}

// serve and start

func (a *Api) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, params := a.router.Find(req.Method, req.URL.Path)
	if handler == nil {
		http.NotFound(w, req)
		return
	}
	handler(&Context{
		W:      w,
		R:      req,
		Params: params,
	})
}

func (api *Api) Start(port string) {
	http.ListenAndServe(port, api)
}
