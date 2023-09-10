package main 

import (
	"log"

	"github.com/Vanv1k/web-course/internal/api"
)

func main() {
	log.Println("Application start!")
	api.StartServer()
	log.Println("Application terminated!")
}
