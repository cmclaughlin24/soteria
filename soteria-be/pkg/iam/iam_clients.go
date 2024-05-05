package iam

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IamHttpClient struct {
	AccessTokenUrl string
	ApiKeyUrl      string
}

func (s *IamHttpClient) VerifyAccessToken(ctx context.Context, token string) (*AccessTokenClaims, error) {
	if token == "" {
		return nil, errors.New("token cannot be an empty string")
	}

	payload, _ := json.Marshal(struct {
		Token string `json:"token"`
	}{token})
	request, err := http.NewRequestWithContext(ctx, "POST", s.AccessTokenUrl, bytes.NewBuffer(payload))

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
		Message string            `json:"message"`
		Data    AccessTokenClaims `json:"data"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data.Data, nil
}

func (s *IamHttpClient) VerifyApiKey(ctx context.Context, key string) (*ApiKeyClaims, error) {
	if key == "" {
		return nil, errors.New("key cannot be an empty string")
	}

	payload, _ := json.Marshal(struct {
		Key string `json:"key"`
	}{key})
	request, err := http.NewRequestWithContext(ctx, "POST", s.ApiKeyUrl, bytes.NewBuffer(payload))

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
		Message string       `json:"message"`
		Data    ApiKeyClaims `json:"data"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data.Data, nil
}

type IamGrpcClient struct {
	client IamServiceClient
}

func NewIamGrpcClient(target string) *IamGrpcClient {
	cc, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		panic(err)
	}

	client := NewIamServiceClient(cc)

	return &IamGrpcClient{client}
}

func (c *IamGrpcClient) VerifyAccessToken(ctx context.Context, token string) (*AccessTokenClaims, error) {
	claims, err := c.client.VerifyAccessToken(ctx, &VerifyTokenRequest{Token: token})

	if err != nil {
		return nil, err
	}

	return &AccessTokenClaims{
		Name:                 claims.GetName(),
		AuthorizationDetails: claims.GetAuthorizationDetails(),
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   claims.GetSub(),
			ExpiresAt: jwt.NewNumericDate(time.Unix(claims.GetExpiresAt(), 0)),
		},
	}, nil
}

func (c *IamGrpcClient) VerifyApiKey(ctx context.Context, key string) (*ApiKeyClaims, error) {
	claims, err := c.client.VerifyApiKey(ctx, &VerifyApiKeyRequest{Key: key})

	if err != nil {
		return nil, err
	}

	return &ApiKeyClaims{
		Sub:                  claims.GetSub(),
		Name:                 claims.GetName(),
		AuthorizationDetails: claims.GetAuthorizationDetails(),
		ExpiresAt:            claims.GetExpiresAt(),
	}, nil
}
