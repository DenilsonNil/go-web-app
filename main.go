package main

import (
	"net/http"
	"webapp/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.GetRoutes()
	http.ListenAndServe(":8080", nil)
}
