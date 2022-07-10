package routers

import "net/http"

// Api Routers
type Router struct {
	Uri      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	ReqAuth  bool
}
