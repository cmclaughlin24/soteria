package ports

import (
	"context"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/domain"
	"github.com/cmclaughlin24/soteria-be/pkg/iam"
)

type ApiKeyService interface {
	Create(context.Context, string, []iam.UserPermission, string) (string, error)
	Remove(context.Context, string) error
	VerifyApiKey(context.Context, string) (*iam.ApiKeyClaims, error)
}

type AuthenticationService interface {
	Signin(context.Context, string, string) (*domain.Tokens, error)
	VerifyAccessToken(context.Context, string) (*iam.AccessTokenClaims, error)
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

type Services struct {
	ApiKey         ApiKeyService
	Authentication AuthenticationService
	Permission     PermissionService
	User           UserService
}
