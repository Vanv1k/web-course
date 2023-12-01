package app

import (
	"log"

	"github.com/Vanv1k/web-course/internal/app/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Vanv1k/web-course/docs"
	_ "github.com/lib/pq"
)

func (a *Application) StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	c := controller.NewController(a.repository)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// r.POST("auth/login", a.Login)

	// r.POST("auth/registration", a.Register)

	// r.GET("auth/logout", a.Logout)

	// r.GET("/consultations", c.GetAllConsultations)
	// r.GET("/consultations/:id", c.GetConsultationByID)
	// r.Use(a.WithAuthCheck(role.Manager, role.Admin)).DELETE("/consultations/delete/:id", c.DeleteConsultation)
	// r.Use(a.WithAuthCheck(role.Manager, role.Admin)).POST("/consultations/create", c.CreateConsultation)
	// r.Use(a.WithAuthCheck(role.Manager, role.Admin)).PUT("/consultations/update/:id", c.UpdateConsultation)
	// r.Use(a.WithAuthCheck(role.Buyer)).POST("/consultations/:id/add-to-request", c.AddConsultationToRequest)
	// r.Use(a.WithAuthCheck(role.Manager, role.Admin)).POST("consultations/:id/addImage", c.AddConsultationImage)
	// r.Use(a.WithAuthCheck(role.Buyer)).GET("/consultations/request/:id", c.GetConsultationsByRequestID)

	// r.Use(a.WithAuthCheck(role.Manager, role.Admin)).GET("/requests", c.GetAllRequests)
	// r.Use(a.WithAuthCheck(role.Buyer)).DELETE("/requests/delete/:id", c.DeleteRequest)
	// r.Use(a.WithAuthCheck(role.Buyer)).PUT("/requests/update/:id", c.UpdateRequest)
	// r.Use(a.WithAuthCheck(role.Buyer)).PUT("/requests/:id/user/update-status", c.UpdateRequestStatusToSendedByUser)
	// r.Use(a.WithAuthCheck(role.Manager, role.Admin)).PUT("/requests/:id/moderator/update-status", c.UpdateRequestStatus)

	// r.Use(a.WithAuthCheck(role.Buyer)).DELETE("/consultation-request/delete/consultation/:id_c/request/:id_r", c.DeleteConsultationRequest)

	// r.Use(a.WithAuthCheck(role.Manager, role.Admin)).GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"Status":  "Ok",
	// 		"Message": "GG",
	// 	})
	// })

	AuthGroup := r.Group("/auth")
	{
		AuthGroup.POST("/registration", a.Register)
		AuthGroup.POST("/login", a.Login)
		AuthGroup.GET("/logout", a.Logout)

	}

	ConsultationGroup := r.Group("/consultations")
	{
		ConsultationGroup.GET("/", c.GetAllConsultations)
		ConsultationGroup.GET("/:id", c.GetConsultationByID)
		ConsultationGroup.GET("/request/:id", c.GetConsultationsByRequestID)
		ConsultationGroup.DELETE("/delete/:id", c.DeleteConsultation)
		ConsultationGroup.PUT("/update/:id", c.UpdateConsultation)
		ConsultationGroup.POST("/create", c.CreateConsultation)
		ConsultationGroup.POST("/:id/add-to-request", c.AddConsultationToRequest)
		ConsultationGroup.POST("/:id/addImage", c.AddConsultationImage)
	}

	RequestGroup := r.Group("/requests")
	{
		RequestGroup.GET("/", c.GetAllRequests)
		RequestGroup.DELETE("/delete/:id", c.DeleteRequest)
		RequestGroup.PUT("/update/:id", c.UpdateRequest)
		RequestGroup.PUT("/:id/user/update-status", c.UpdateRequestStatusToSendedByUser)
		RequestGroup.PUT("/:id/moderator/update-status", c.UpdateRequestStatus)
	}

	ConsultationRequestGroup := r.Group("/consultation-request")
	{
		ConsultationRequestGroup.DELETE("/delete/consultation/:id_c/request/:id_r", c.DeleteConsultationRequest)
	}

	r.Run()

	log.Println("Server down")
}
