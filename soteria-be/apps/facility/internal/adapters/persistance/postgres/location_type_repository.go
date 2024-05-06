package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/domain"
)

type LocationTypeRepository struct {
	db *sql.DB
}

func NewLocationTypeRepository(db *sql.DB) *LocationTypeRepository {
	return &LocationTypeRepository{
		db: db,
	}
}

func (r *LocationTypeRepository) FindAll(ctx context.Context) ([]domain.LocationType, error) {
	query := "SELECT id, name, enableChildren, createdBy, updatedBy FROM locationType"

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	locationTypes := make([]domain.LocationType, 0)

	for rows.Next() {
		var lt domain.LocationType
		var updatedBy sql.NullString

		err = rows.Scan(
			&lt.Id,
			&lt.Name,
			&lt.EnableChildren,
			&lt.CreatedBy,
			&updatedBy,
		)

		if err != nil {
			return nil, err
		}

		lt.UpdatedBy = updatedBy.String
		locationTypes = append(locationTypes, lt)
	}

	return locationTypes, nil
}

func (r *LocationTypeRepository) FindOne(ctx context.Context, id int) (*domain.LocationType, error) {
	query := "SELECT id, name, enableChildren, createdBy, updatedBy FROM locationType WHERE id = $1 "

	row := r.db.QueryRowContext(ctx, query, id)
	var lt domain.LocationType
	var updatedBy sql.NullString

	err := row.Scan(
		&lt.Id,
		&lt.Name,
		&lt.EnableChildren,
		&lt.CreatedBy,
		&updatedBy,
	)

	if err != nil {
		return nil, err
	}

	lt.UpdatedBy = updatedBy.String

	return &lt, nil
}

func (r *LocationTypeRepository) Create(ctx context.Context, lt domain.LocationType) (*domain.LocationType, error) {
	stmt, err := r.db.PrepareContext(ctx, `
		INSERT INTO locationType(
			name,
			enableChildren,
			createdBy
		)
		VALUES($1, $2, $3)
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, lt.Name, lt.EnableChildren, lt.CreatedBy)

	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	lt.Id = int(id)

	return &lt, nil
}

func (r *LocationTypeRepository) Update(ctx context.Context, lt domain.LocationType) (*domain.LocationType, error) {
	stmt, err := r.db.PrepareContext(ctx, "UPDATE locationType SET name = $2, enableChildren = $3, updatedBy = $4, updatedAt = $5 WHERE id = $1")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, lt.Id, lt.Name, lt.EnableChildren, lt.UpdatedBy, time.Now())

	if err != nil {
		return nil, err
	}

	return &lt, nil
}

func (r *LocationTypeRepository) Remove(ctx context.Context, id int) error {
	stmt, err := r.db.PrepareContext(ctx, "DELETE FROM locationType WHERE id = $1")

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, id); err != nil {
		return err
	}

	return nil
}
