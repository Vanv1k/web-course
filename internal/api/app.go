package app

import (
	"time"

	"github.com/Vanv1k/web-course/internal/app/controller"
	"github.com/Vanv1k/web-course/internal/app/dsn"
	"github.com/Vanv1k/web-course/internal/app/redis"
	"github.com/Vanv1k/web-course/internal/app/repository"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Application struct {
	config     *Config
	repository *repository.Repository
	redis      *redis.Client
	controller *controller.Controller
}

type Config struct {
	JWT struct {
		Token         string
		SigningMethod jwt.SigningMethod
		ExpiresIn     time.Duration
	}
}

func New() (*Application, error) {
	_ = godotenv.Load()

	config := &Config{}
	err := envconfig.Process("", config)
	if err != nil {
		return nil, err
	}

	repo, err := repository.New(dsn.SetConnectionString())
	if err != nil {
		return nil, err
	}

	redisClient, err := redis.New()
	if err != nil {
		return nil, err
	}

	c := controller.NewController(repo)

	return &Application{config: config, repository: repo, redis: redisClient, controller: c}, nil
}
