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
		FacilityRepository: NewFacilityRepository(db),
	}, nil
}

func createTables(db *sql.DB) error {
	createFacilityTable := `
		CREATE TABLE IF NOT EXISTS facilities(
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

	return nil
}
