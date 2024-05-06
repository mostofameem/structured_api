package web

import (
	"log"
)

func StartServer() {
	RegisterRoutes()
	log.Println("Server started on :3000")
}
