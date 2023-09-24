package main

import (
	"log"

	app "github.com/Vanv1k/web-course/internal/api"
)

func main() {
	log.Println("Application start!")

	application, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	application.StartServer()

	log.Println("Application terminated!")
}
