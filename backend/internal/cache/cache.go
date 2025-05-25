package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/likoscp/Advanced-Programming-2/backend/internal/config"
)

func NewRedisClient(config *config.ConfigRedis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: config.ADDR + ":" + config.PORT,
		DB:   0,
	})
}
