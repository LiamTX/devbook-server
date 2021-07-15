package main

import (
	"api/src/config"
	"api/src/router"
	"log"
	"net/http"
)

func main() {
	config.Load()
	routes := router.Generate()

	print("server started\n")
	log.Fatal(http.ListenAndServe(":3000", routes))
}
