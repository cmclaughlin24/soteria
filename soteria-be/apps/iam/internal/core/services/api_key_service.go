package services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/domain"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/pkg/auth"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/pkg/hash"
	"github.com/google/uuid"
)

const maxApiKeyBytes = 72
const apiKeyServiceCacheKey = "api-key-service"

type ApiKeyService struct {
	repository  ports.ApiKeyRepository
	hashService hash.HashService
	cache       ports.CacheManager
}

func NewApiKeyService(repository ports.ApiKeyRepository, hashService hash.HashService, cache ports.CacheManager) *ApiKeyService {
	return &ApiKeyService{
		repository:  repository,
		hashService: hashService,
		cache:       cache,
	}
}

func (s *ApiKeyService) Create(ctx context.Context, name string, permissions []domain.UserPermission, createdBy string) (string, error) {
	apiKeyId := uuid.New()
	expiresAt := time.Now().AddDate(1, 0, 0)
	claims := domain.ApiKeyClaims{
		Sub:                  apiKeyId.String(),
		Name:                 name,
		AuthorizationDetails: auth.PackPermissions(permissions), // Todo: Should probably load permissions exist since not coming from database.
		ExpiresAt:            expiresAt.Unix(),
	}
	encodedKey, err := s.generateApiKey(&claims)

	if err != nil {
		return "", err
	}

	compressedKey := s.compressApiKey(encodedKey)
	hashedKey, err := s.hashService.Hash(compressedKey)

	if err != nil {
		return "", err
	}

	_, err = s.repository.Create(ctx, domain.ApiKey{
		Id:        apiKeyId.String(),
		Name:      name,
		ApiKey:    hashedKey,
		ExpiresAt: expiresAt,
		CreatedBy: createdBy,
	})

	if err != nil {
		return "", err
	}

	return encodedKey, nil
}

func (s *ApiKeyService) Remove(ctx context.Context, id string) error {
	return s.repository.Remove(ctx, id)
}

/*
Yields a struct containing the api key claims if the token is valid.
*/
func (s *ApiKeyService) VerifyApiKey(ctx context.Context, key string) (*domain.ApiKeyClaims, error) {
	claims, err := s.extractApiKeyData(key)

	if err != nil {
		return nil, err
	}

	apiKey, err := s.getApiKey(ctx, claims.Sub)

	if err != nil {
		return nil, err
	}

	compressedKey := s.compressApiKey(key)

	if err := s.hashService.Compare(compressedKey, apiKey.ApiKey); err != nil {
		return nil, err
	}

	if s.isExpired(apiKey.ExpiresAt) {
		log.Println(apiKey.ExpiresAt)
		return nil, errors.New("api key has expired")
	}

	return claims, nil
}

func (s *ApiKeyService) getApiKey(ctx context.Context, id string) (*domain.ApiKey, error) {
	var apiKey *domain.ApiKey
	cacheKey := cachKeyHash(apiKeyServiceCacheKey, id)

	if err := s.cache.Get(ctx, cacheKey, &apiKey); err != nil || apiKey == nil {
		// Fixme: Add informational log that the cache failed to find a result or had an error.
		apiKey, err = s.repository.FindOne(ctx, id)

		if err != nil || apiKey == nil {
			return nil, fmt.Errorf("api key id=%s could not be retreived from cache or database", id)
		}

		s.cache.Set(ctx, cacheKey, apiKey)
	}

	return apiKey, nil
}

func (s ApiKeyService) generateApiKey(claims *domain.ApiKeyClaims) (string, error) {
	out, err := json.Marshal(claims)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString([]byte(out)), nil
}

func (s ApiKeyService) extractApiKeyData(key string) (*domain.ApiKeyClaims, error) {
	data, err := base64.StdEncoding.DecodeString(key)

	if err != nil {
		return nil, err
	}

	var claims domain.ApiKeyClaims

	if err := json.Unmarshal(data, &claims); err != nil {
		return nil, err
	}

	return &claims, nil
}

func (s ApiKeyService) compressApiKey(key string) string {
	keyBytes := []byte(key)

	if !s.exceedsMaxBytes(keyBytes) {
		return key
	}

	return string(keyBytes[len(keyBytes)-maxApiKeyBytes:])
}

func (s ApiKeyService) exceedsMaxBytes(data []byte) bool {
	return len([]byte(data)) > maxApiKeyBytes
}

/*
Yields true if the current datetime is greater than the expiration
time.
*/
func (s ApiKeyService) isExpired(expiresAt time.Time) bool {
	return time.Now().After(expiresAt)
}
