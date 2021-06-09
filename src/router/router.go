package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Returns router with routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
