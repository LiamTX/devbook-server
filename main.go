package main

import (
	"api/src/router"
	"log"
	"net/http"
)

func main() {
	routes := router.Generate()

	print("server started\n")
	log.Fatal(http.ListenAndServe(":3000", routes))
}
