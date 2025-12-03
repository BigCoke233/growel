package growel

import (
	"net/http"
)

// create API engine

func New() *Engine {
	var a Engine
	a.router = NewRouter()
	return &a
}

// add routes

func (api *Engine) GET(path string, h Handler) {
	api.router.Add("GET", path, h)
}

func (api *Engine) POST(path string, h Handler) {
	api.router.Add("POST", path, h)
}

func (api *Engine) PUT(path string, h Handler) {
	api.router.Add("PUT", path, h)
}

func (api *Engine) DELETE(path string, h Handler) {
	api.router.Add("DELETE", path, h)
}

// serve and start

func (a *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, params := a.router.Find(req.Method, req.URL.Path)
	if handler == nil {
		http.NotFound(w, req)
		return
	}

	err := req.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	handler(&Context{
		W:      w,
		R:      req,
		Params: params,
		Querys: req.URL.Query(),
		Form:   req.Form,
	})
}

func (api *Engine) Start(port string) {
	http.ListenAndServe(port, api)
}
