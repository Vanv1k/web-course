package controller

import (
	"github.com/Vanv1k/web-course/internal/app/repository"
	"github.com/gin-gonic/gin"
)

func Login(repository *repository.Repository, c *gin.Context) {
	repository.Login()
}

func Register(repository *repository.Repository, c *gin.Context) {
	repository.Register()
}

func Logout(repository *repository.Repository, c *gin.Context) {
	repository.Logout()
}
