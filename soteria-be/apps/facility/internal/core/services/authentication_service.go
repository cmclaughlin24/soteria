package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/cmclaughlin24/soteria-be/pkg/iam"
)

type AuthenticationService struct {
}

func NewAuthenticationService() *AuthenticationService {
	return &AuthenticationService{}
}

func (s *AuthenticationService) VerifyAccessToken(ctx context.Context, token string) (*iam.AccessTokenClaims, error) {
	if token == "" {
		return nil, errors.New("token cannot be an empty string")
	}

	payload, _ := json.Marshal(struct {
		Token string `json:"token"`
	}{token})
	// Todo: Properly handle iam service url and remove hardcoded value.
	url := "http://iam/authentication/verify"

	request, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected statusCode=%d but received statusCode=%d", http.StatusOK, response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var data struct {
		Message string                `json:"message"`
		Data    iam.AccessTokenClaims `json:"data"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data.Data, nil
}

func (s *AuthenticationService) VerifyApiKey(ctx context.Context, key string) (*iam.ApiKeyClaims, error) {
	return nil, errors.New("AuthenticiationService VerifyApiKey not implemented")
}
