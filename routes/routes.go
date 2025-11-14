package routes

import (
	"net/http"
	"webapp/controllers"
)

func GetRoutes() {
	http.HandleFunc("/", controllers.Index)
}
