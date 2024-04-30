package domain

import (
	"github.com/golang-jwt/jwt/v5"
)

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func NewTokens(accessToken, refreshToken string) *Tokens {
	return &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

type RefreshTokenClaims struct {
	RefreshTokenId string `json:"refreshTokenId"`
	jwt.RegisteredClaims
}
