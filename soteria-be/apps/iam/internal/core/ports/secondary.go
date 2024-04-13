package ports

import (
	"context"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/domain"
)

type ApiKeyRepository interface {
	FindOne(context.Context, string) (*domain.ApiKey, error)
	Create(context.Context, domain.ApiKey) (*domain.ApiKey, error)
	Remove(context.Context, string) error
}

type CacheManager interface {
	Get(context.Context, string, any) error
	Set(context.Context, string, any) error
	Del(context.Context, string) error
}

type PermissionRepository interface {
	FindAll(context.Context) ([]domain.Permission, error)
	FindOne(context.Context, string) (*domain.Permission, error)
	Create(context.Context, domain.Permission) (*domain.Permission, error)
	Update(context.Context, domain.Permission) (*domain.Permission, error)
	Remove(context.Context, string) error
}

type UserRepository interface {
	FindAll(context.Context) ([]domain.User, error)
	FindOne(context.Context, string) (*domain.User, error)
	FindByEmail(context.Context, string) (*domain.User, error)
	Create(context.Context, domain.User) (*domain.User, error)
	Update(context.Context, domain.User) (*domain.User, error)
	Remove(context.Context, string) error
}

type Repositories struct {
	ApiKeyRepository     ApiKeyRepository
	PermissionRepository PermissionRepository
	UserRepository       UserRepository
}
