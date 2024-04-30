package services

import (
	"context"
	"errors"
	"time"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/domain"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/pkg/hash"
	"github.com/cmclaughlin24/soteria-be/pkg/iam"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtSignOptions struct {
	Secret     string
	Audience   string
	Issuer     string
	Ttl        int64
	RefreshTtl int64
}

type AuthenticationService struct {
	repository   ports.UserRepository
	tokenStorage *TokenStorage
	hashService  hash.HashService
	signOptions  JwtSignOptions
}

func NewAuthenticationService(
	repository ports.UserRepository,
	tokenStorage *TokenStorage,
	hashService hash.HashService,
	signOptions JwtSignOptions,
) *AuthenticationService {
	// Todo: Add validation on JwtSignOptions.
	return &AuthenticationService{
		repository:   repository,
		tokenStorage: tokenStorage,
		hashService:  hashService,
		signOptions:  signOptions,
	}
}

/*
Yields a struct containing the access and refresh tokens.
*/
func (s *AuthenticationService) Signin(ctx context.Context, email, password string) (*domain.Tokens, error) {
	u, err := s.repository.FindByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if err := s.hashService.Compare(password, u.Password); err != nil {
		return nil, err
	}

	return s.generateTokens(u)
}

/*
Yields a struct containing the access token claims if the token is valid.

Note: An access token id (jti) is validated against the jti stored to ensure
only a single access token is issued for a user.
*/
func (s *AuthenticationService) VerifyAccessToken(ctx context.Context, token string) (*iam.AccessTokenClaims, error) {
	claims, err := s.verify(token, &iam.AccessTokenClaims{})

	if err != nil {
		return nil, err
	}

	userId, err := claims.GetSubject()

	if err != nil {
		return nil, err
	}

	jwtId := claims.(*iam.AccessTokenClaims).ID

	if isValid, err := s.tokenStorage.ValidateAccess(ctx, userId, jwtId); !isValid || err != nil {
		return nil, errors.New("failed to validate access token id")
	}

	return claims.(*iam.AccessTokenClaims), nil
}

/*
Yields a struct containing the new access and refresh tokens if the token is valid.

Note: A token is invalidated when it is redeemed for the first time (replay attack)
or if an attempt to redeem a previously used token is made (revoke token family).
This is known as Refresh Token Rotation.
*/
func (s *AuthenticationService) RefreshAccessToken(ctx context.Context, token string) (*domain.Tokens, error) {
	claims, err := s.verify(token, &domain.RefreshTokenClaims{})

	if err != nil {
		return nil, err
	}

	userId, err := claims.GetSubject()

	if err != nil {
		return nil, err
	}

	refreshTokenId := claims.(*domain.RefreshTokenClaims).RefreshTokenId

	if isValid, err := s.tokenStorage.ValidateRefresh(ctx, userId, refreshTokenId); !isValid || err != nil {
		if err := s.tokenStorage.Remove(ctx, userId); err != nil {
			// Fixme: Add log message indicating the refresh token could not be invalidated.
		}

		return nil, errors.New("failed to validate refresh token id")
	}

	u, err := s.repository.FindOne(ctx, userId)

	if err != nil {
		return nil, err
	}

	return s.generateTokens(u)
}

/*
Yields a struct containing JWT access and refresh tokens.
*/
func (s *AuthenticationService) generateTokens(u *domain.User) (*domain.Tokens, error) {
	jwtId := uuid.New()
	refreshTokenId := uuid.New()
	permissions := iam.PackPermissions(u.Permissions)

	accessToken, err := s.signToken(u.Id, s.signOptions.Ttl, jwt.MapClaims{
		"name":                  u.Name,
		"authorization_details": permissions,
		"jti":                   jwtId.String(),
	})

	if err != nil {
		return nil, err
	}

	refreshToken, err := s.signToken(u.Id, s.signOptions.RefreshTtl, jwt.MapClaims{
		"refreshTokenId": refreshTokenId.String(),
	})

	if err != nil {
		return nil, err
	}

	store := TokenStore{JwtId: jwtId.String(), RefreshTokenId: refreshTokenId.String()}

	if err := s.tokenStorage.Insert(context.TODO(), u.Id, store); err != nil {
		return nil, err
	}

	return domain.NewTokens(accessToken, refreshToken), nil
}

/*
Yields a JSON Web Token (JWT).
*/
func (s *AuthenticationService) signToken(userId string, expiresIn int64, claims jwt.MapClaims) (string, error) {
	claims["sub"] = userId
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Unix() + expiresIn
	claims["aud"] = s.signOptions.Audience
	claims["iss"] = s.signOptions.Issuer

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.signOptions.Secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthenticationService) verify(token string, claims jwt.Claims) (jwt.Claims, error) {
	t, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.signOptions.Secret), nil
	}, jwt.WithAudience(s.signOptions.Audience), jwt.WithIssuer(s.signOptions.Issuer))

	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, errors.New("invalid token")
	}

	return t.Claims, nil
}
