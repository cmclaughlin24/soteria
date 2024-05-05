package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/domain"
)

type PgxPermissionRepository struct {
	db *sql.DB
}

func NewPermissionRepository(connection *sql.DB) *PgxPermissionRepository {
	return &PgxPermissionRepository{
		db: connection,
	}
}

func (r *PgxPermissionRepository) FindAll(ctx context.Context) ([]domain.Permission, error) {
	query := "SELECT id, resource, action FROM permission"

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	permissions := make([]domain.Permission, 0)

	for rows.Next() {
		var p domain.Permission

		err = rows.Scan(
			&p.Id,
			&p.Resource,
			&p.Action,
		)

		if err != nil {
			return nil, err
		}

		permissions = append(permissions, p)
	}

	return permissions, nil
}

func (r *PgxPermissionRepository) FindOne(ctx context.Context, id string) (*domain.Permission, error) {
	query := "SELECT id, resource, action FROM permission WHERE id = $1"

	row := r.db.QueryRowContext(ctx, query, id)
	var p domain.Permission

	err := row.Scan(
		&p.Id,
		&p.Resource,
		&p.Action,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PgxPermissionRepository) Create(ctx context.Context, p domain.Permission) (*domain.Permission, error) {
	stmt, err := r.db.PrepareContext(ctx, "INSERT INTO permission(resource, action) VALUES ($1, $2)")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, p.Resource, p.Action)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PgxPermissionRepository) Update(ctx context.Context, p domain.Permission) (*domain.Permission, error) {
	stmt, err := r.db.PrepareContext(ctx, "UPDATE permission SET resource = $2, action = $3, updatedAt = $4 WHERE id = $1")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, p.Id, p.Resource, p.Action, time.Now())

	if err != nil {
		return nil, err
	}

	return &p, err
}

func (r *PgxPermissionRepository) Remove(ctx context.Context, id string) error {
	stmt, err := r.db.PrepareContext(ctx, "DELETE FROM permission WHERE id = $1")

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, id); err != nil {
		return err
	}

	return nil
}
