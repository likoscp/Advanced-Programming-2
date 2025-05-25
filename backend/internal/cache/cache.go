package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/likoscp/Advanced-Programming-2/backend/internal/config"
)

func NewRedisClient(config *config.Config) *redis.Client {
    return redis.NewClient(&redis.Options{
        Addr: "redis:" + config.ConfigRedis.ADDR,
        DB:   0,
    })
}