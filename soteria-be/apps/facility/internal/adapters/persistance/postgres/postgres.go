package postgres

import (
	"database/sql"
	"errors"
	"os"

	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/ports"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() (*ports.Repositories, error) {
	connectionString := os.Getenv("DB_CONNECTION_STRING")

	if connectionString == "" {
		return nil, errors.New("environment DB_CONNECTION_STRING cannot be an empty string")
	}

	db, err := sql.Open("pgx", connectionString)

	if err != nil {
		return nil, err
	}

	if err := createTables(db); err != nil {
		return nil, err
	}

	return &ports.Repositories{
		Facility:     NewFacilityRepository(db),
		Location:     NewLocationRepository(db),
		LocationType: NewLocationTypeRepository(db),
	}, nil
}

func createTables(db *sql.DB) error {
	createFacilityTable := `
		CREATE TABLE IF NOT EXISTS facility(
			code VARCHAR(25) NOT NULL,
			name VARCHAR(250) NOT NULL,
			createdBy VARCHAR(100) NOT NULL,
			createdAt TIMESTAMP DEFAULT NOW(),
			updatedBy VARCHAR(100) NULL,
			updatedAt TIMESTAMP NULL,
			PRIMARY KEY(code)
		)
	`

	if _, err := db.Exec(createFacilityTable); err != nil {
		return err
	}

	createLocationTable := `
		CREATE TABLE IF NOT EXISTS location(
			id SERIAL PRIMARY KEY,
			code VARCHAR(25) NOT NULL,
			name VARCHAR(250) NOT NULL,
			createdBy VARCHAR(100) NOT NULL,
			createdAt TIMESTAMP DEFAULT NOW(),
			updatedBy VARCHAR(100) NULL,
			updatedAt TIMESTAMP NULL,
			facilityCode VARCHAR(25) NOT NULL,
			parentId INTEGER NULL,
			FOREIGN KEY(facilityCode) REFERENCES facility(code) ON DELETE CASCADE,
			FOREIGN KEY(parentId) REFERENCES location(id) ON DELETE CASCADE,
			CONSTRAINT Uq_Facility_Location UNIQUE NULLS NOT DISTINCT (facilityCode, parentId, code)
		)
	`

	if _, err := db.Exec(createLocationTable); err != nil {
		return err
	}

	createLocationTypeTable := `
		CREATE TABLE IF NOT EXISTS locationType(
			id SERIAL PRIMARY KEY,
			name VARCHAR(250) NOT NULL,
			enableChildren BOOLEAN DEFAULT FALSE,
			createdBy VARCHAR(100) NOT NULL,
			createdAt TIMESTAMP DEFAULT NOW(),
			updatedBy VARCHAR(100) NULL,
			updatedAt TIMESTAMP NULL
		)
	`

	if _, err := db.Exec(createLocationTypeTable); err != nil {
		return err
	}

	return nil
}
