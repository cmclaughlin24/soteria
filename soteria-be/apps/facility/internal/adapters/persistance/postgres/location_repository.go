package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/domain"
)

type LocationRepository struct {
	db *sql.DB
}

func NewLocationRepository(db *sql.DB) *LocationRepository {
	return &LocationRepository{
		db: db,
	}
}

func (r *LocationRepository) FindAll(ctx context.Context) ([]domain.Location, error) {
	query := "SELECT id, code, name, createdBy, updatedBy, facilityCode, parentId FROM location"

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	locations := make([]domain.Location, 0)

	for rows.Next() {
		var l domain.Location
		var updatedBy sql.NullString
		var parentId sql.NullInt64

		err = rows.Scan(
			&l.Id,
			&l.Code,
			&l.Name,
			&l.CreatedBy,
			&updatedBy,
			&l.FacilityCode,
			&parentId,
		)

		if err != nil {
			return nil, err
		}

		l.UpdatedBy = updatedBy.String
		l.ParentId = int(parentId.Int64)

		locations = append(locations, l)
	}

	return locations, nil
}

func (r *LocationRepository) FindOne(ctx context.Context, id int) (*domain.Location, error) {
	query := "SELECT id, code, name, createdBy, updatedBy, facilityCode, parentId FROM location WHERE id = $1"

	row := r.db.QueryRowContext(ctx, query, id)
	var l domain.Location
	var updatedBy sql.NullString
	var parentId sql.NullInt64

	err := row.Scan(
		&l.Id,
		&l.Code,
		&l.Name,
		&l.CreatedBy,
		&updatedBy,
		&l.FacilityCode,
		&parentId,
	)

	if err != nil {
		return nil, err
	}

	l.UpdatedBy = updatedBy.String
	l.ParentId = int(parentId.Int64)

	return &l, nil
}

func (r *LocationRepository) Create(ctx context.Context, l domain.Location) (*domain.Location, error) {
	stmt, err := r.db.PrepareContext(ctx, `
		INSERT INTO location(
				code,
				name,
				facilityCode,
				parentId,
				createdBy
			)
		VALUES($1, $2, $3, $4, $5)
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, l.Code, l.Name, l.FacilityCode, newNullInt64(l.ParentId), l.CreatedBy)

	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	l.Id = int(id)

	return &l, nil
}

func (r *LocationRepository) Update(ctx context.Context, l domain.Location) (*domain.Location, error) {
	stmt, err := r.db.PrepareContext(ctx, "UPDATE location SET code = $2, name = $3, parentId = $4, updatedAt = $5 WHERE id = $1")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, l.Id, l.Code, l.Name, newNullInt64(l.ParentId), time.Now())

	if err != nil {
		return nil, err
	}

	return &l, nil
}

func (r *LocationRepository) Remove(ctx context.Context, id int) error {
	stmt, err := r.db.PrepareContext(ctx, "DELETE FROM location WHERE id = $1")

	if err != nil {
		return nil
	}

	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, id); err != nil {
		return err
	}

	return nil
}
