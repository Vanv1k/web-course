package app

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Vanv1k/web-course/internal/app/controller"
	"github.com/Vanv1k/web-course/internal/app/ds"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func (a *Application) StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.Static("/styles", "./resources/styles")
	r.Static("/js", "./resources/js")
	r.Static("/img", "./resources/img")
	r.Static("/hacker", "./resources")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		var consultations []ds.Consultation
		consultations, err := a.repository.GetAllConsultations()
		if err != nil { // если не получилось
			log.Printf("cant get product by id %v", err)
			return
		}
		searchQuery := c.DefaultQuery("fsearch", "")

		if searchQuery == "" {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"services": consultations,
			})
			return
		}

		var result []ds.Consultation

		for _, consultation := range consultations {
			if strings.Contains(strings.ToLower(consultation.Name), strings.ToLower(searchQuery)) {
				result = append(result, consultation)
			}
		}

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"services":    result,
			"search_text": searchQuery,
		})
	})

	r.POST("/delete/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			// Обработка ошибки
			log.Printf("cant get consultation by id %v", err)
			c.Redirect(http.StatusMovedPermanently, "/")
		}
		a.repository.DeleteConsultation(id)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.GET("/service/:id", func(c *gin.Context) {
		var consultation *ds.Consultation

		id, err := strconv.Atoi(c.Param("id"))
		consultation, err = a.repository.GetConsultationByID(uint(id))
		if err != nil {
			// Обработка ошибки
			log.Printf("cant get service by id %v", err)
			return
		}

		c.HTML(http.StatusOK, "card.tmpl", consultation)
	})

	//------------------------------------------------------------------------------
	r.GET("/consultations", func(c *gin.Context) {
		controller.GetAllConsultations(a.repository, c)
	})
	r.GET("/consultations/:id", func(c *gin.Context) {
		controller.GetConsultationByID(a.repository, c)
	})
	r.DELETE("/consultations/delete/:id", func(c *gin.Context) {
		controller.DeleteConsultation(a.repository, c)
	})
	r.POST("/consultations/create", func(c *gin.Context) {
		controller.CreateConsultation(a.repository, c)
	})
	r.PUT("/consultations/update/:id", func(c *gin.Context) {
		controller.UpdateConsultation(a.repository, c)
	})
	r.POST("/consultations/:id/add-to-request", func(c *gin.Context) {
		controller.AddConsultationToRequest(a.repository, c)
	})

	r.GET("/requests", func(c *gin.Context) {
		controller.GetAllRequests(a.repository, c)
	})
	r.GET("/requests/:id", func(c *gin.Context) {
		controller.GetConsultationsByRequestID(a.repository, c)
	})
	r.DELETE("/requests/delete/:id", func(c *gin.Context) {
		controller.DeleteRequest(a.repository, c)
	})
	r.PUT("/requests/update/:id", func(c *gin.Context) {
		controller.UpdateRequest(a.repository, c)
	})
	r.PUT("/requests/user/:id/update-status", func(c *gin.Context) {
		controller.UpdateRequestStatus(a.repository, c)
	})
	r.PUT("/requests/moderator/:id/update-status", func(c *gin.Context) {
		controller.UpdateRequestStatus(a.repository, c)
	})

	r.DELETE("/consultation-request/delete/consultation/:id_c/request/:id_r", func(c *gin.Context) {
		controller.DeleteConsultationRequest(a.repository, c)
	})

	r.Run()

	log.Println("Server down")
}
