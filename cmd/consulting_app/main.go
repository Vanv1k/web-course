package main

import (
	"log"

	app "github.com/Vanv1k/web-course/internal/api"
)

// @title IT Services
// @version 1.0
// @description Consultation app

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	log.Println("Application start!")

	application, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	application.StartServer()

	log.Println("Application terminated!")
}
