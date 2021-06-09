package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		uri:      "/users",
		method:   http.MethodPost,
		function: controllers.Create,
		auth:     false,
	},
	{
		uri:      "/users",
		method:   http.MethodGet,
		function: controllers.Find,
		auth:     false,
	},
	{
		uri:      "/users/{id}",
		method:   http.MethodGet,
		function: controllers.FindOne,
		auth:     false,
	},
	{
		uri:      "/users/{id}",
		method:   http.MethodPut,
		function: controllers.Update,
		auth:     false,
	},
	{
		uri:      "/users/{id}",
		method:   http.MethodDelete,
		function: controllers.Delete,
		auth:     false,
	},
}
