package cache

import (
	"fmt"
	"os"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/adapters/cache/memory"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/adapters/cache/redis"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
)

func Connect() (ports.CacheManager, error) {
	cacheType := os.Getenv("CACHE_TYPE")

	switch cacheType {
	case "redis":
		return redis.Connect()
	case "memory":
		return memory.NewMemoryCacheManager(), nil
	default:
		return nil, fmt.Errorf("invalid cache type: %s is not implemented", cacheType)
	}
}
