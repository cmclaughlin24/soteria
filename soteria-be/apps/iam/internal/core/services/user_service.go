package services

import (
	"context"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/domain"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/pkg/hash"
)

type UserService struct {
	repository   ports.UserRepository
	hashService  hash.HashService
	tokenStorage *TokenStorage
}

func NewUserService(repository ports.UserRepository, hashService hash.HashService, tokenStorage *TokenStorage) *UserService {
	return &UserService{
		repository:   repository,
		hashService:  hashService,
		tokenStorage: tokenStorage,
	}
}

func (s *UserService) FindAll(ctx context.Context) ([]domain.User, error) {
	return s.repository.FindAll(ctx)
}

func (s *UserService) FindOne(ctx context.Context, id string) (*domain.User, error) {
	return s.repository.FindOne(ctx, id)
}

func (s *UserService) Create(ctx context.Context, user domain.User) (*domain.User, error) {
	hashedPassword, err := s.hashService.Hash(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	return s.repository.Create(ctx, user)
}

/*
Update a user.

Note: If a user is updated, any active issued JWT will be invalidate to
prevent revoked access to resources.
*/
func (s *UserService) Update(ctx context.Context, user domain.User) (*domain.User, error) {
	updated, err := s.repository.Update(ctx, user)

	if err != nil {
		return nil, err
	}

	if err := s.tokenStorage.Remove(ctx, updated.Id); err != nil {
		// Fixme: Add a log message indicating the tokens could not be invalidated.
	}

	return updated, nil
}

/*
Removes a user.

Note: If a user is removed, any active issued JWT will be invalidate to
prevent revoked access to resources.
*/
func (s *UserService) Remove(ctx context.Context, id string) error {
	if err := s.repository.Remove(ctx, id); err != nil {
		return err
	}

	if err := s.tokenStorage.Remove(ctx, id); err != nil {
		// Fixme: Add log message indicating the tokens could not be invalidated.
	}

	return nil
}
