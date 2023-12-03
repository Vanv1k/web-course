package app

import (
	"log"

	"github.com/Vanv1k/web-course/internal/app/role"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Vanv1k/web-course/docs"
	_ "github.com/lib/pq"
)

func (a *Application) StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	AuthGroup := r.Group("/auth")
	{
		AuthGroup.POST("/registration", a.Register)
		AuthGroup.POST("/login", a.Login)
		AuthGroup.GET("/logout", a.Logout)

	}

	ConsultationGroup := r.Group("/consultations")
	{
		ConsultationGroup.GET("/", a.controller.GetAllConsultations)
		ConsultationGroup.GET("/:id", a.controller.GetConsultationByID)
		ConsultationGroup.Use(a.WithAuthCheck(role.Buyer)).GET("/request", a.controller.GetConsultationsByRequestID)
		ConsultationGroup.Use(a.WithAuthCheck(role.Manager, role.Admin)).DELETE("/delete/:id", a.controller.DeleteConsultation)
		ConsultationGroup.Use(a.WithAuthCheck(role.Manager, role.Admin)).PUT("/update/:id", a.controller.UpdateConsultation)
		ConsultationGroup.Use(a.WithAuthCheck(role.Manager, role.Admin)).POST("/create", a.controller.CreateConsultation)
		ConsultationGroup.Use(a.WithAuthCheck(role.Buyer)).POST("/:id/add-to-request", a.controller.AddConsultationToRequest)
		ConsultationGroup.Use(a.WithAuthCheck(role.Manager, role.Admin)).POST("/:id/addImage", a.controller.AddConsultationImage)
	}

	RequestGroup := r.Group("/requests")
	{
		RequestGroup.Use(a.WithAuthCheck(role.Buyer, role.Manager, role.Admin)).GET("/", a.controller.GetAllRequests)
		RequestGroup.Use(a.WithAuthCheck(role.Manager, role.Admin)).PUT("/:id/moderator/update-status", a.controller.UpdateRequestStatus)
		RequestGroup.Use(a.WithAuthCheck(role.Buyer)).DELETE("/delete/:id", a.controller.DeleteRequest)
		RequestGroup.Use(a.WithAuthCheck(role.Buyer)).PUT("/update/:id", a.controller.UpdateRequest)
		RequestGroup.Use(a.WithAuthCheck(role.Buyer)).PUT("/:id/user/update-status", a.controller.UpdateRequestStatusToSendedByUser)
	}

	ConsultationRequestGroup := r.Group("/consultation-request")
	{
		ConsultationRequestGroup.Use(a.WithAuthCheck(role.Buyer)).DELETE("/delete/consultation/:id_c/request/:id_r", a.controller.DeleteConsultationRequest)
	}

	r.Run()

	log.Println("Server down")
}
