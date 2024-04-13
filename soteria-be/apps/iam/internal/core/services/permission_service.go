package services

import (
	"context"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/domain"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
)

type PermissionService struct {
	repository ports.PermissionRepository
}

func NewPermissionService(repository ports.PermissionRepository) *PermissionService {
	return &PermissionService{
		repository: repository,
	}
}

func (s *PermissionService) FindAll(ctx context.Context) ([]domain.Permission, error) {
	return s.repository.FindAll(ctx)
}

func (s *PermissionService) FindOne(ctx context.Context, id string) (*domain.Permission, error) {
	return s.repository.FindOne(ctx, id)
}

func (s *PermissionService) Create(ctx context.Context, permission domain.Permission) (*domain.Permission, error) {
	return s.repository.Create(ctx, permission)
}

func (s *PermissionService) Update(ctx context.Context, permission domain.Permission) (*domain.Permission, error) {
	return s.repository.Update(ctx, permission)
}

func (s *PermissionService) Remove(ctx context.Context, id string) error {
	return s.repository.Remove(ctx, id)
}
