package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRoute = Route{
	uri:      "/login",
	method:   http.MethodPost,
	function: controllers.Login,
	auth:     false,
}
