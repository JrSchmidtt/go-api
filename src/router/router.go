package router

import "github.com/gorilla/mux"

// Build Router return a new Router instance.
func Build() *mux.Router{
	return mux.NewRouter()
}