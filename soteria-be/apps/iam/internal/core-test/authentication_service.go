package coretest

import (
	"context"
	"errors"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/domain"
)

type SuccessAuthenticationService struct {
}

func NewSuccessAuthenticationService() *SuccessAuthenticationService {
	return &SuccessAuthenticationService{}
}

func (s *SuccessAuthenticationService) Signin(_ context.Context, _, _ string) (*domain.Tokens, error) {
	return &domain.Tokens{}, nil
}

func (s *SuccessAuthenticationService) VerifyAccessToken(_ context.Context, _ string) (*domain.AccessTokenClaims, error) {
	return &domain.AccessTokenClaims{}, nil
}

func (s *SuccessAuthenticationService) RefreshAccessToken(_ context.Context, _ string) (*domain.Tokens, error) {
	return &domain.Tokens{}, nil
}

type ErrorAuthenticationService struct{}

func NewErrorAuthenticationService() *ErrorAuthenticationService {
	return &ErrorAuthenticationService{}
}

func (s *ErrorAuthenticationService) Signin(_ context.Context, _, _ string) (*domain.Tokens, error) {
	return nil, errors.New("mock error from coretest package")
}

func (s *ErrorAuthenticationService) VerifyAccessToken(_ context.Context, _ string) (*domain.AccessTokenClaims, error) {
	return nil, errors.New("mock error from coretest package")
}

func (s *ErrorAuthenticationService) RefreshAccessToken(_ context.Context, _ string) (*domain.Tokens, error) {
	return nil, errors.New("mock error from coretest package")
}
