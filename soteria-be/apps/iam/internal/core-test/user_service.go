package coretest

import (
	"context"
	"errors"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/domain"
)

type SuccessUserService struct {
}

func NewSuccessUserService() *SuccessUserService {
	return &SuccessUserService{}
}

func (s *SuccessUserService) FindAll(_ context.Context) ([]domain.User, error) {
	return make([]domain.User, 0), nil
}

func (s *SuccessUserService) FindOne(_ context.Context, _ string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (s *SuccessUserService) Create(_ context.Context, _ domain.User) (*domain.User, error) {
	return &domain.User{}, nil
}

func (s *SuccessUserService) Update(_ context.Context, _ domain.User) (*domain.User, error) {
	return &domain.User{}, nil
}

func (s *SuccessUserService) Remove(_ context.Context, _ string) error {
	return nil
}

type ErrorUserService struct{}

func NewErrorUserService() *ErrorUserService {
	return &ErrorUserService{}
}

func (s *ErrorUserService) FindAll(_ context.Context) ([]domain.User, error) {
	return nil, errors.New("mock error from coretest package")
}

func (s *ErrorUserService) FindOne(_ context.Context, _ string) (*domain.User, error) {
	return nil, errors.New("mock error from coretest package")
}

func (s *ErrorUserService) Create(_ context.Context, _ domain.User) (*domain.User, error) {
	return nil, errors.New("mock error from coretest package")
}

func (s *ErrorUserService) Update(_ context.Context, _ domain.User) (*domain.User, error) {
	return nil, errors.New("mock error from coretest package")
}

func (s *ErrorUserService) Remove(_ context.Context, _ string) error {
	return errors.New("mock error from coretest package")
}
