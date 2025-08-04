package db

// redis.go
// This file contains the Redis client setup for the application.

import (
	"fmt"

	"angi.id/internal/shared/config"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Acfg.Redis.Host, config.Acfg.Redis.Port),
		Password: config.Acfg.Redis.Password, // no password set
		DB:       config.Acfg.Redis.DB,
	})
	return rdb
}
