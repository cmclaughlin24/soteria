package domain

import (
	"time"
)

type ApiKey struct {
	Id        string
	Name      string
	ApiKey    string
	ExpiresAt time.Time
	CreatedBy string
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
