package services

import (
	"context"

	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/domain"
	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/ports"
)

type LocationService struct {
	repository ports.LocationRepository
}

func NewLocationService(repository ports.LocationRepository) *LocationService {
	return &LocationService{
		repository: repository,
	}
}

func (s *LocationService) FindAll(ctx context.Context) ([]domain.Location, error) {
	// Fixme: Add support for accepting options to specify if the child locatoins should be returned.
	return s.repository.FindAll(ctx)
}

func (s *LocationService) FindOne(ctx context.Context, id int) (*domain.Location, error) {
	// Fixme: Add support for accepting options to specify if the child locations should be returned.
	return s.repository.FindOne(ctx, id)
}

func (s *LocationService) Create(ctx context.Context, location domain.Location) (*domain.Location, error) {
	return s.repository.Create(ctx, location)
}

func (s *LocationService) Update(ctx context.Context, location domain.Location) (*domain.Location, error) {
	return s.repository.Update(ctx, location)
}

func (s *LocationService) Remove(ctx context.Context, int int) error {
	return s.repository.Remove(ctx, int)
}
