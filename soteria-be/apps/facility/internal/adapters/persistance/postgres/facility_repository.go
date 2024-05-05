package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/domain"
)

type FacilityRepository struct {
	db *sql.DB
}

func NewFacilityRepository(db *sql.DB) *FacilityRepository {
	return &FacilityRepository{
		db: db,
	}
}

func (r *FacilityRepository) FindAll(ctx context.Context) ([]domain.Facility, error) {
	query := "SELECT code, name, createdBy, updatedBy FROM facility"

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	facilities := make([]domain.Facility, 0)

	for rows.Next() {
		var f domain.Facility
		var updatedBy sql.NullString

		err = rows.Scan(
			&f.Code,
			&f.Name,
			&f.CreatedBy,
			&updatedBy,
		)

		if err != nil {
			return nil, err
		}

		f.UpdatedBy = updatedBy.String
		facilities = append(facilities, f)
	}

	return facilities, nil
}

func (r *FacilityRepository) FindOne(ctx context.Context, code string) (*domain.Facility, error) {
	query := "SELECT code, name, createdBy, updatedBy FROM facility WHERE code = $1"

	row := r.db.QueryRowContext(ctx, query, code)
	var f domain.Facility
	var updatedBy sql.NullString

	err := row.Scan(
		&f.Code,
		&f.Name,
		&f.CreatedBy,
		&updatedBy,
	)

	if err != nil {
		return nil, err
	}

	f.UpdatedBy = updatedBy.String

	return &f, nil
}

func (r *FacilityRepository) Create(ctx context.Context, f domain.Facility) (*domain.Facility, error) {
	stmt, err := r.db.PrepareContext(ctx, "INSERT INTO facility(code, name, createdBy) VALUES($1, $2, $3)")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, f.Code, f.Name, f.CreatedBy)

	if err != nil {
		return nil, err
	}

	return &f, nil
}

func (r *FacilityRepository) Update(ctx context.Context, f domain.Facility) (*domain.Facility, error) {
	stmt, err := r.db.PrepareContext(ctx, "UPDATE facility SET name = $2, updatedBy = $3, updatedAt = $4 WHERE code = $1")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, f.Code, f.Name, f.UpdatedBy, time.Now())

	if err != nil {
		return nil, err
	}

	return &f, nil
}

func (r *FacilityRepository) Remove(ctx context.Context, code string) error {
	stmt, err := r.db.PrepareContext(ctx, "DELETE FROM facility WHERE code = $1")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, code)

	if err != nil {
		return err
	}

	return nil
}
