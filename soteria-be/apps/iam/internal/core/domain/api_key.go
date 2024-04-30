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
