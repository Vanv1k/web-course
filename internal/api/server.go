package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	log.Println("Server start up")
	
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,  gin.H{
			"message": "hello",
		})
	})
	r.Run()

	log.Println("Server down")
}
