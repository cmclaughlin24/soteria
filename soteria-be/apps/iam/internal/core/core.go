package core

import (
	"os"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/adapters/cache"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/adapters/persistance"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/services"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/pkg/hash"
)

func Init() (*ports.Drivers, error) {
	repositories, err := persistance.Connect()

	if err != nil {
		return nil, err
	}

	cacheManager, err := cache.Connect()

	if err != nil {
		return nil, err
	}

	tokenStorage := services.NewTokenStorage(cacheManager)

	return &ports.Drivers{
		ApiKeyService: services.NewApiKeyService(repositories.ApiKeyRepository, hash.BcryptService{}, cacheManager),
		AuthenticationService: services.NewAuthenticationService(repositories.UserRepository, tokenStorage, hash.BcryptService{}, services.JwtSignOptions{
			Secret:     os.Getenv("JWT_SECRET"),
			Audience:   os.Getenv("JWT_TOKEN_AUDIENCE"),
			Issuer:     os.Getenv("JWT_TOKEN_ISSUER"),
			Ttl:        3600,
			RefreshTtl: 86400,
		}),
		PermissionsService: services.NewPermissionService(repositories.PermissionRepository),
		UserService:        services.NewUserService(repositories.UserRepository, hash.BcryptService{}, tokenStorage),
	}, nil
}
