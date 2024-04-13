package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/domain"
)

type PgxApiKeyRepository struct {
	db *sql.DB
}

func NewApiKeyRepository(connection *sql.DB) *PgxApiKeyRepository {
	return &PgxApiKeyRepository{
		db: connection,
	}
}

func (r *PgxApiKeyRepository) FindOne(ctx context.Context, id string) (*domain.ApiKey, error) {
	query := `
		SELECT id,
			name,
			apiKey,
			expiresAt
		FROM api_key
		WHERE id = $1 AND deletedAt IS NULL
	`

	var key domain.ApiKey
	row := r.db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&key.Id,
		&key.Name,
		&key.ApiKey,
		&key.ExpiresAt,
	)

	if err != nil {
		return nil, err
	}

	return &key, nil
}

func (r *PgxApiKeyRepository) Create(ctx context.Context, key domain.ApiKey) (*domain.ApiKey, error) {
	stmt, err := r.db.PrepareContext(ctx, `
		INSERT INTO api_key(
			id,
			name,
			apiKey,
			expiresAt,
			createdBy
		)
		VALUES ($1, $2, $3, $4, $5)
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		key.Id,
		key.Name,
		key.ApiKey,
		key.ExpiresAt,
		key.CreatedBy,
	)

	if err != nil {
		return nil, err
	}

	return &key, nil
}

func (r *PgxApiKeyRepository) Remove(ctx context.Context, id string) error {
	stmt, err := r.db.PrepareContext(ctx, "UPDATE api_key SET deletedAt = $2 WHERE id = $1")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id, time.Now())

	if err != nil {
		return err
	}

	return nil
}
