package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
)

const tokenStorageCacheKey = "token-storage"

type TokenStore struct {
	JwtId          string `json:"jwtId"`
	RefreshTokenId string `json:"refreshTokenId"`
}

type TokenStorage struct {
	cache ports.CacheManager
}

func NewTokenStorage(cache ports.CacheManager) *TokenStorage {
	return &TokenStorage{
		cache: cache,
	}
}

func (s *TokenStorage) Insert(ctx context.Context, userId string, store TokenStore) error {
	cacheKey := s.generateCacheKey(userId)
	return s.cache.Set(ctx, cacheKey, store)
}

func (s *TokenStorage) ValidateAccess(ctx context.Context, userId string, tokenId string) (bool, error) {
	var store TokenStore
	cacheKey := s.generateCacheKey(userId)

	if err := s.cache.Get(ctx, cacheKey, &store); err != nil {
		return false, err
	}

	return store.JwtId == tokenId, nil
}

func (s *TokenStorage) ValidateRefresh(ctx context.Context, userId string, tokenId string) (bool, error) {
	var store TokenStore
	cacheKey := s.generateCacheKey(userId)

	if err := s.cache.Get(ctx, cacheKey, &store); err != nil {
		return false, err
	}

	return store.RefreshTokenId == tokenId, nil
}

func (s *TokenStorage) Remove(ctx context.Context, userId string) error {
	cacheKey := s.generateCacheKey(userId)
	return s.cache.Del(context.TODO(), cacheKey)
}

func (s *TokenStorage) generateCacheKey(userId string) string {
	return cachKeyHash(tokenStorageCacheKey, userId)
}

func cachKeyHash(key string, args ...any) string {
	if len(args) == 0 {
		return key
	}

	cacheKey := fmt.Sprintf("%s::", key)
	stringArgs := make([]string, 0, len(args))

	for _, arg := range args {
		out, err := json.Marshal(arg)

		if err != nil {
			// Todo: Improve error handling in cachKeyHash function.
			log.Println(err)
		}

		stringArgs = append(stringArgs, string(out[:]))
	}

	cacheKey += strings.Join(stringArgs, ",")

	return cacheKey
}
