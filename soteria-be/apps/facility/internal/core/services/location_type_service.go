package services

import (
	"context"

	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/domain"
	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/ports"
)

type LocationTypeService struct {
	repository ports.LocationTypeRepository
}

func NewLocationTypeService(repository ports.LocationTypeRepository) *LocationTypeService {
	return &LocationTypeService{
		repository: repository,
	}
}

func (s *LocationTypeService) FindAll(ctx context.Context) ([]domain.LocationType, error) {
	return s.repository.FindAll(ctx)
}

func (s *LocationTypeService) FindOne(ctx context.Context, id int) (*domain.LocationType, error) {
	return s.repository.FindOne(ctx, id)
}

func (s *LocationTypeService) Create(ctx context.Context, lt domain.LocationType) (*domain.LocationType, error) {
	return s.repository.Create(ctx, lt)
}

func (s *LocationTypeService) Update(ctx context.Context, lt domain.LocationType) (*domain.LocationType, error) {
	return s.repository.Update(ctx, lt)
}

func (s *LocationTypeService) Remove(ctx context.Context, id int) error {
	return s.repository.Remove(ctx, id)
}
