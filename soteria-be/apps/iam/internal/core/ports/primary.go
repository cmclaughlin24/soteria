package ports

import (
	"context"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/domain"
)

type ApiKeyService interface {
	Create(context.Context, string, []domain.UserPermission, string) (string, error)
	Remove(context.Context, string) error
	VerifyApiKey(context.Context, string) (*domain.ApiKeyClaims, error)
}

type AuthenticationService interface {
	Signin(context.Context, string, string) (*domain.Tokens, error)
	VerifyAccessToken(context.Context, string) (*domain.AccessTokenClaims, error)
	RefreshAccessToken(context.Context, string) (*domain.Tokens, error)
}

type PermissionService interface {
	FindAll(context.Context) ([]domain.Permission, error)
	FindOne(context.Context, string) (*domain.Permission, error)
	Create(context.Context, domain.Permission) (*domain.Permission, error)
	Update(context.Context, domain.Permission) (*domain.Permission, error)
	Remove(context.Context, string) error
}

type UserService interface {
	FindAll(context.Context) ([]domain.User, error)
	FindOne(context.Context, string) (*domain.User, error)
	Create(context.Context, domain.User) (*domain.User, error)
	Update(context.Context, domain.User) (*domain.User, error)
	Remove(context.Context, string) error
}

type Drivers struct {
	ApiKeyService         ApiKeyService
	AuthenticationService AuthenticationService
	PermissionsService    PermissionService
	UserService           UserService
}
