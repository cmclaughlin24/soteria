package services

import (
	"context"

	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/domain"
	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/ports"
)

type FacilityService struct {
	repository ports.FacilityRepository
}

func NewFacilityService(repository ports.FacilityRepository) *FacilityService {
	return &FacilityService{
		repository: repository,
	}
}

func (s *FacilityService) FindAll(ctx context.Context) ([]domain.Facility, error) {
	return s.repository.FindAll(ctx)
}

func (s *FacilityService) FindOne(ctx context.Context, code string) (*domain.Facility, error) {
	return s.repository.FindOne(ctx, code)
}

func (s *FacilityService) Create(ctx context.Context, facility domain.Facility) (*domain.Facility, error) {
	return s.repository.Create(ctx, facility)
}

func (s *FacilityService) Update(ctx context.Context, facility domain.Facility) (*domain.Facility, error) {
	return s.repository.Update(ctx, facility)
}

func (s *FacilityService) Remove(ctx context.Context, code string) error {
	return s.repository.Remove(ctx, code)
}
