package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

const servicePrefix = "awesome_service." // наш префикс сервиса

type RedisConfig struct {
	Host        string
	Password    string
	Port        int
	User        string
	DialTimeout time.Duration
	ReadTimeout time.Duration
}

type Client struct {
	cfg    RedisConfig
	client *redis.Client
}

func New() (*Client, error) {
	client := &Client{}

	client.cfg.Host = "0.0.0.0"
	client.cfg.Password = "password"
	client.cfg.Port = 6379
	client.cfg.User = ""
	client.cfg.DialTimeout = 10
	client.cfg.ReadTimeout = 10

	redisClient := redis.NewClient(&redis.Options{
		Password:    client.cfg.Password,
		Username:    client.cfg.User,
		Addr:        client.cfg.Host + ":" + strconv.Itoa(client.cfg.Port),
		DB:          0,
		DialTimeout: client.cfg.DialTimeout,
		ReadTimeout: client.cfg.ReadTimeout,
	})

	client.client = redisClient

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("can't ping redis: %w", err)
	}

	return client, nil
}

func (c *Client) Close() error {
	return c.client.Close()
}
