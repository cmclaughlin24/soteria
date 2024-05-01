package ports

import (
	"context"

	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/domain"
	"github.com/cmclaughlin24/soteria-be/pkg/iam"
)

type AuthenticationService interface {
	VerifyAccessToken(context.Context, string) (*iam.AccessTokenClaims, error)
	VerifyApiKey(context.Context, string) (*iam.ApiKeyClaims, error)
}

type FacilityService interface {
	FindAll(context.Context) ([]domain.Facility, error)
	FindOne(context.Context, string) (*domain.Facility, error)
	Create(context.Context, domain.Facility) (*domain.Facility, error)
	Update(context.Context, domain.Facility) (*domain.Facility, error)
	Remove(context.Context, string) error
}

type Services struct {
	Authentication AuthenticationService
	Facility       FacilityService
}
