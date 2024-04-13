package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func Connect() (*RedisCacheManager, error) {
	redisMode := os.Getenv("REDIS_MODE")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	var client *redis.Client

	switch redisMode {
	case "cluster":
	case "sentinel":
	default:
		client = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
		})
	}

	if status := client.Ping(context.TODO()); status.Err() != nil {
		return nil, status.Err()
	}

	return NewRedisCacheManager(client), nil
}
