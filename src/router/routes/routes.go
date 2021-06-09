package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes type
type Route struct {
	uri      string
	method   string
	function func(http.ResponseWriter, *http.Request)
	auth     bool
}

// Config all routes in router
func Config(r *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		r.HandleFunc(route.uri, route.function).Methods(route.method)
	}

	return r
}
