package app

import (
	"github.com/Vanv1k/web-course/internal/app/dsn"
	"github.com/Vanv1k/web-course/internal/app/repository"
	"github.com/joho/godotenv"
)

type Application struct {
	repository *repository.Repository
}

func New() (Application, error) {
	_ = godotenv.Load()
	repo, err := repository.New(dsn.SetConnectionString())
	if err != nil {
		return Application{}, err
	}

	return Application{repository: repo}, nil
}
