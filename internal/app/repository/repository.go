package repository

import (
	"log"

	minioclient "github.com/Vanv1k/web-course/internal/minioClient"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db          *gorm.DB
	minioClient *minioclient.MinioClient
}

func New(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	minioClient, err := minioclient.NewMinioClient()
	if err != nil {
		log.Println("error here start!")
		return nil, err
	}

	return &Repository{
		db:          db,
		minioClient: minioClient,
	}, nil
}
