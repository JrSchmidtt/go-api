package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Api Routers
type Router struct {
	Uri      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	ReqAuth  bool
}

//Put all routes inside the Router
func ConfigRouter(r *mux.Router) *mux.Router{
	router := userRouter
	for _, route := range router {
		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}
	return r
}
