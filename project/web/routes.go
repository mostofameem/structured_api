package web

import (
	"net/http"
	"project/web/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/create", handlers.Create)
	http.HandleFunc("/list", handlers.List)
}
