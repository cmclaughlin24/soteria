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

type LocationService interface {
	FindAll(context.Context) ([]domain.Location, error)
	FindOne(context.Context, int) (*domain.Location, error)
	Create(context.Context, domain.Location) (*domain.Location, error)
	Update(context.Context, domain.Location) (*domain.Location, error)
	Remove(context.Context, int) error
}

type LocationTypeService interface {
	FindAll(context.Context) ([]domain.LocationType, error)
	FindOne(context.Context, int) (*domain.LocationType, error)
	Create(context.Context, domain.LocationType) (*domain.LocationType, error)
	Update(context.Context, domain.LocationType) (*domain.LocationType, error)
	Remove(context.Context, int) error
}

type Services struct {
	Authentication AuthenticationService
	Facility       FacilityService
	Location       LocationService
	LocationType   LocationTypeService
}
