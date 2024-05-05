package core

import (
	"os"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/adapters/cache"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/adapters/persistance"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/services"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/pkg/hash"
)

func Init() (*ports.Services, error) {
	repositories, err := persistance.Connect()

	if err != nil {
		return nil, err
	}

	cacheManager, err := cache.Connect()

	if err != nil {
		return nil, err
	}

	tokenStorage := services.NewTokenStorage(cacheManager)

	return &ports.Services{
		ApiKey: services.NewApiKeyService(repositories.ApiKey, hash.BcryptService{}, cacheManager),
		Authentication: services.NewAuthenticationService(repositories.User, tokenStorage, hash.BcryptService{}, services.JwtSignOptions{
			Secret:     os.Getenv("JWT_SECRET"),
			Audience:   os.Getenv("JWT_TOKEN_AUDIENCE"),
			Issuer:     os.Getenv("JWT_TOKEN_ISSUER"),
			Ttl:        3600,
			RefreshTtl: 86400,
		}),
		Permission: services.NewPermissionService(repositories.Permission),
		User:       services.NewUserService(repositories.User, hash.BcryptService{}, tokenStorage),
	}, nil
}
