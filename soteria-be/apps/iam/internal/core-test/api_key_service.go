package coretest

import (
	"context"
	"errors"

	"github.com/cmclaughlin24/soteria-be/pkg/iam"
)

type SuccessApiKeyService struct {
}

func NewSuccessApiKeyService() *SuccessApiKeyService {
	return &SuccessApiKeyService{}
}

func (s *SuccessApiKeyService) Create(_ context.Context, _ string, _ []iam.UserPermission, _ string) (string, error) {
	return "", nil
}

func (s *SuccessApiKeyService) Remove(_ context.Context, _ string) error {
	return nil
}

func (s *SuccessApiKeyService) VerifyApiKey(_ context.Context, _ string) (*iam.ApiKeyClaims, error) {
	return &iam.ApiKeyClaims{}, nil
}

type ErrorApiKeyService struct{}

func NewErrorApiKeyService() *ErrorApiKeyService {
	return &ErrorApiKeyService{}
}

func (s *ErrorApiKeyService) Create(_ context.Context, _ string, _ []iam.UserPermission, _ string) (string, error) {
	return "", errors.New("mock error from coretest package")
}

func (s *ErrorApiKeyService) Remove(_ context.Context, _ string) error {
	return errors.New("mock error from coretest package")
}

func (s *ErrorApiKeyService) VerifyApiKey(_ context.Context, _ string) (*iam.ApiKeyClaims, error) {
	return nil, errors.New("mock error from coretest package")
}
