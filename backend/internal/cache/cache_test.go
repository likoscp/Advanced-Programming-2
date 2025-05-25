package cache

import (
	"context"
	"testing"
	"time"

	"github.com/likoscp/Advanced-Programming-2/backend/internal/config"
)

func TestRedis(t *testing.T) {
	config := config.ConfigRedis{
		ADDR: "localhost",
		PORT: "6379",
	}

	r := NewRedisClient(&config)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	s, err := r.Ping(ctx).Result()

	if err != nil {
		t.Fatalf("cannot connect to redis, err:%v, s:%s", err, s)
	}

}
