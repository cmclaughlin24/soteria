package coretest

import (
	"context"
	"errors"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/domain"
)

type SuccessPermissionService struct {
}

func NewSuccessPermissionService() *SuccessPermissionService {
	return &SuccessPermissionService{}
}

func (s *SuccessPermissionService) FindAll(_ context.Context) ([]domain.Permission, error) {
	return make([]domain.Permission, 0), nil
}

func (s *SuccessPermissionService) FindOne(_ context.Context, _ string) (*domain.Permission, error) {
	return &domain.Permission{}, nil
}

func (s *SuccessPermissionService) Create(_ context.Context, _ domain.Permission) (*domain.Permission, error) {
	return &domain.Permission{}, nil
}

func (s *SuccessPermissionService) Update(_ context.Context, _ domain.Permission) (*domain.Permission, error) {
	return &domain.Permission{}, nil
}

func (s *SuccessPermissionService) Remove(_ context.Context, _ string) error {
	return nil
}

type ErrorPermissionService struct{}

func NewErrorPermissionService() *ErrorPermissionService {
	return &ErrorPermissionService{}
}

func (s *ErrorPermissionService) FindAll(_ context.Context) ([]domain.Permission, error) {
	return nil, errors.New("mock error from coretest package")
}

func (s *ErrorPermissionService) FindOne(_ context.Context, _ string) (*domain.Permission, error) {
	return nil, errors.New("mock error from coretest package")
}

func (s *ErrorPermissionService) Create(_ context.Context, _ domain.Permission) (*domain.Permission, error) {
	return nil, errors.New("mock error from coretest package")
}

func (s *ErrorPermissionService) Update(_ context.Context, _ domain.Permission) (*domain.Permission, error) {
	return nil, errors.New("mock error from coretest package")
}

func (s *ErrorPermissionService) Remove(_ context.Context, _ string) error {
	return errors.New("mock error from coretest package")
}
