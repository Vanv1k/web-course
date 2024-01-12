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
		AuthGroup.GET("/logout", a.WithAuthCheck(role.Buyer, role.Manager, role.Admin), a.Logout)

	}

	ConsultationGroup := r.Group("/consultations")
	{
		ConsultationGroup.GET("/:id", a.controller.GetConsultationByID)
		ConsultationGroup.GET("/", a.WithAuthCheck(role.Buyer, role.Manager, role.Admin), a.controller.GetAllConsultations)
		ConsultationGroup.GET("/request/:id", a.WithAuthCheck(role.Buyer, role.Manager, role.Admin), a.controller.GetConsultationsByRequestID)
		ConsultationGroup.DELETE("/delete/:id", a.WithAuthCheck(role.Manager, role.Admin), a.controller.DeleteConsultation)
		ConsultationGroup.PUT("/update/:id", a.WithAuthCheck(role.Manager, role.Admin), a.controller.UpdateConsultation)
		ConsultationGroup.POST("/create", a.WithAuthCheck(role.Manager, role.Admin), a.controller.CreateConsultation)
		ConsultationGroup.POST("/:id/add-to-request", a.WithAuthCheck(role.Buyer, role.Manager, role.Admin), a.controller.AddConsultationToRequest)
		ConsultationGroup.POST("/:id/addImage", a.WithAuthCheck(role.Manager, role.Admin), a.controller.AddConsultationImage)
	}

	RequestGroup := r.Group("/requests")
	{
		RequestGroup.GET("/", a.WithAuthCheck(role.Buyer, role.Manager, role.Admin), a.controller.GetAllRequests)
		RequestGroup.PUT("/:id/moderator/update-status", a.WithAuthCheck(role.Manager, role.Admin), a.controller.UpdateRequestStatus)
		RequestGroup.DELETE("/delete/:id", a.WithAuthCheck(role.Buyer, role.Manager, role.Admin), a.controller.DeleteRequest)
		RequestGroup.PUT("/update/:id", a.WithAuthCheck(role.Buyer, role.Manager, role.Admin), a.controller.UpdateRequest)
		RequestGroup.PUT("/:id/user/update-status", a.WithAuthCheck(role.Buyer, role.Manager, role.Admin), a.controller.UpdateRequestStatusToSendedByUser)
	}

	ConsultationRequestGroup := r.Group("/consultation-request")
	{
		ConsultationRequestGroup.Use(a.WithAuthCheck(role.Buyer, role.Manager, role.Admin)).DELETE("/delete/consultation/:id_c/request/:id_r", a.controller.DeleteConsultationRequest)
	}

	r.Run()

	log.Println("Server down")
}
