package memory

import (
	"context"
	"encoding/json"
	"fmt"
)

type MemoryCacheManager struct {
	cache map[string]string
}

func NewMemoryCacheManager() *MemoryCacheManager {
	return &MemoryCacheManager{
		cache: make(map[string]string),
	}
}

func (manager *MemoryCacheManager) Get(_ context.Context, key string, data any) error {
	val, ok := manager.cache[key]

	if !ok {
		return fmt.Errorf("invalid key: %s not found in cache", key)
	}

	if val == "" {
		return nil
	}

	return json.Unmarshal([]byte(val), &data)
}

func (manager *MemoryCacheManager) Set(_ context.Context, key string, data any) error {
	out, err := json.Marshal(data)

	if err != nil {
		return nil
	}

	manager.cache[key] = string(out[:])

	return nil
}

func (manager *MemoryCacheManager) Del(_ context.Context, key string) error {
	if _, ok := manager.cache[key]; !ok {
		return fmt.Errorf("invalid key: %s not found in cache", key)
	}

	delete(manager.cache, key)

	return nil
}
