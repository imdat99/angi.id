package db

// redis.go
// This file contains the Redis client setup for the application.

import (
	"fmt"

	s "angi.id/internal/shared"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", s.Acfg.Redis.Host, s.Acfg.Redis.Port),
		Password: s.Acfg.Redis.Password, // no password set
		DB:       s.Acfg.Redis.DB,
	})
	return rdb
}
