package coretest

import (
	"context"
	"errors"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/domain"
	"github.com/cmclaughlin24/soteria-be/pkg/auth"
)

type SuccessApiKeyService struct {
}

func NewSuccessApiKeyService() *SuccessApiKeyService {
	return &SuccessApiKeyService{}
}

func (s *SuccessApiKeyService) Create(_ context.Context, _ string, _ []auth.UserPermission, _ string) (string, error) {
	return "", nil
}

func (s *SuccessApiKeyService) Remove(_ context.Context, _ string) error {
	return nil
}

func (s *SuccessApiKeyService) VerifyApiKey(_ context.Context, _ string) (*domain.ApiKeyClaims, error) {
	return &domain.ApiKeyClaims{}, nil
}

type ErrorApiKeyService struct{}

func NewErrorApiKeyService() *ErrorApiKeyService {
	return &ErrorApiKeyService{}
}

func (s *ErrorApiKeyService) Create(_ context.Context, _ string, _ []auth.UserPermission, _ string) (string, error) {
	return "", errors.New("mock error from coretest package")
}

func (s *ErrorApiKeyService) Remove(_ context.Context, _ string) error {
	return errors.New("mock error from coretest package")
}

func (s *ErrorApiKeyService) VerifyApiKey(_ context.Context, _ string) (*domain.ApiKeyClaims, error) {
	return nil, errors.New("mock error from coretest package")
}
