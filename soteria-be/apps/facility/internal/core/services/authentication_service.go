package services

import (
	"context"

	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/ports"
	"github.com/cmclaughlin24/soteria-be/pkg/iam"
)

type AuthenticationService struct {
	client ports.IamClient
}

func NewAuthenticationService(client ports.IamClient) *AuthenticationService {
	return &AuthenticationService{
		client: client,
	}
}

func (s *AuthenticationService) VerifyAccessToken(ctx context.Context, token string) (*iam.AccessTokenClaims, error) {
	return s.client.VerifyAccessToken(ctx, token)
}

func (s *AuthenticationService) VerifyApiKey(ctx context.Context, key string) (*iam.ApiKeyClaims, error) {
	return s.client.VerifyApiKey(ctx, key)
}
