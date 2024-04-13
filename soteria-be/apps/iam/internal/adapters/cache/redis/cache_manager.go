package redis

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type RedisCacheManager struct {
	client *redis.Client
}

func NewRedisCacheManager(client *redis.Client) *RedisCacheManager {
	return &RedisCacheManager{client: client}
}

func (manager *RedisCacheManager) Get(ctx context.Context, key string, data any) error {
	val, err := manager.client.Get(ctx, key).Result()

	if err != nil {
		return err
	}

	if val == "" {
		return nil
	}

	return json.Unmarshal([]byte(val), &data)
}

func (manager *RedisCacheManager) Set(ctx context.Context, key string, data any) error {
	out, err := json.Marshal(data)

	if err != nil {
		return err
	}

	status := manager.client.Set(ctx, key, out, 0)

	if status.Err() != nil {
		return status.Err()
	}

	return nil
}

func (manager *RedisCacheManager) Del(ctx context.Context, key string) error {
	cmd := manager.client.Del(ctx, key)
	return cmd.Err()
}
