package iam

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

const ClaimsContextKey = "claims"

type Claims interface {
	GetSubject() (string, error)
	GetAuthorizationDetails() []string
}

func SetContext(ctx context.Context, claims Claims) context.Context {
	return context.WithValue(ctx, ClaimsContextKey, claims)
}

func ClaimsFromContext(ctx context.Context) Claims {
	return ctx.Value(ClaimsContextKey).(Claims)
}

type AccessTokenClaims struct {
	Name                 string   `json:"name"`
	AuthorizationDetails []string `json:"authorization_details"`
	jwt.RegisteredClaims
}

func (c AccessTokenClaims) GetAuthorizationDetails() []string {
	return c.AuthorizationDetails
}

type ApiKeyClaims struct {
	Sub                  string   `json:"sub"`
	Name                 string   `json:"name"`
	AuthorizationDetails []string `json:"authorization_details"`
	ExpiresAt            int64    `json:"exp"`
}

func (c ApiKeyClaims) GetSubject() (string, error) {
	return c.Sub, nil
}

func (c ApiKeyClaims) GetAuthorizationDetails() []string {
	return c.AuthorizationDetails
}
