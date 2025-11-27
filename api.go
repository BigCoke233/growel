package growel

import (
	"net/http"
)

type Api struct {
	router *Router
}

func New() *Api {
	var a Api
	a.router = NewRouter()
	return &a
}

func (api *Api) Start(port string) {
	http.ListenAndServe(port, api.router)
}

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
