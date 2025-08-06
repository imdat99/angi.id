package container

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

// singleton transient scoped
type Container struct {
	DbPool      *sql.DB
	RedisClient *redis.Client
}

func NewContainer(dbPool *sql.DB, redisClient *redis.Client) *Container {
	return &Container{
		DbPool:      dbPool,
		RedisClient: redisClient,
	}
}
