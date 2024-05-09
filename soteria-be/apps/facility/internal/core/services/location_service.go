package services

import (
	"context"
	"fmt"

	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/domain"
	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/ports"
)

type LocationService struct {
	locationRepository     ports.LocationRepository
	locationTypeRepository ports.LocationTypeRepository
}

func NewLocationService(locationRepository ports.LocationRepository, locationTypeRepository ports.LocationTypeRepository) *LocationService {
	return &LocationService{
		locationRepository:     locationRepository,
		locationTypeRepository: locationTypeRepository,
	}
}

func (s *LocationService) FindAll(ctx context.Context) ([]domain.Location, error) {
	// Fixme: Add support for accepting options to specify if the child locatoins should be returned.
	return s.locationRepository.FindAll(ctx)
}

func (s *LocationService) FindOne(ctx context.Context, id int) (*domain.Location, error) {
	// Fixme: Add support for accepting options to specify if the child locations should be returned.
	return s.locationRepository.FindOne(ctx, id)
}

func (s *LocationService) Create(ctx context.Context, location domain.Location) (*domain.Location, error) {
	if err := s.isChildrenEnabled(ctx, location.ParentId); err != nil {
		return nil, err
	}

	return s.locationRepository.Create(ctx, location)
}

func (s *LocationService) Update(ctx context.Context, location domain.Location) (*domain.Location, error) {
	if err := s.isChildrenEnabled(ctx, location.ParentId); err != nil {
		return nil, err
	}

	return s.locationRepository.Update(ctx, location)
}

func (s *LocationService) Remove(ctx context.Context, int int) error {
	return s.locationRepository.Remove(ctx, int)
}

func (s *LocationService) isChildrenEnabled(ctx context.Context, id int) error {
	if id == 0 {
		return nil
	}

	// Todo: It would make sense to retrieve the parent location w/the location type in
	//       in one transaction instead of executing seperate queries.
	parent, err := s.locationRepository.FindOne(ctx, id)

	if err != nil {
		return err
	}

	locationType, err := s.locationTypeRepository.FindOne(ctx, parent.LocationTypeId)

	if err != nil {
		return err
	}

	if !locationType.EnableChildren {
		return fmt.Errorf(
			"parent location %s is of type %s and cannot have children",
			parent.Name,
			locationType.Name,
		)
	}

	return nil
}
