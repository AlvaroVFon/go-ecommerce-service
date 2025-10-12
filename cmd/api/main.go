package main

import (
	"log"
	"net/http"

	"ecommerce-service/internal/bootstrap"
)

func main() {
	boot, err := bootstrap.Bootstrap()
	if err != nil {
		log.Fatal("Failed to bootstrap application:", err)
	}

	err = http.ListenAndServe(boot.Config.AppHost+":"+boot.Config.AppPort, boot.Router)
	log.Println("Server running on " + boot.Config.AppHost + ":" + boot.Config.AppPort)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
