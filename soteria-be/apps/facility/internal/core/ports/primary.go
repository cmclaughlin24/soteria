package ports

import (
	"context"

	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/domain"
)

type FacilityService interface {
	FindAll(context.Context) ([]domain.Facility, error)
	FindOne(context.Context, string) (*domain.Facility, error)
	Create(context.Context, domain.Facility) (*domain.Facility, error)
	Update(context.Context, domain.Facility) (*domain.Facility, error)
	Remove(context.Context, string) error
}
