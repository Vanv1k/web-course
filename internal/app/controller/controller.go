package controller

import (
	"github.com/Vanv1k/web-course/internal/app/repository"
)

type Controller struct {
	Repo *repository.Repository
}

func NewController(repo *repository.Repository) *Controller {
	return &Controller{Repo: repo}
}
