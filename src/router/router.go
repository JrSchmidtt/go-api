package router

import (
	"api/src/router/routers"

	"github.com/gorilla/mux"
)

//Return a new Router instance.
func Build() *mux.Router{
	router := mux.NewRouter()
	return routers.ConfigRouter(router)
}