package postgres

import (
	"database/sql"
	"errors"
	"os"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
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
		ApiKeyRepository:     NewApiKeyRepository(db),
		PermissionRepository: NewPermissionRepository(db),
		UserRepository:       NewUserRepository(db),
	}, nil
}

func createTables(db *sql.DB) error {
	createPermissionsTable := `
		CREATE TABLE IF NOT EXISTS permission(
			id uuid DEFAULT gen_random_uuid(),
			resource VARCHAR(100) NOT NULL,
			action VARCHAR(100) NOT NULL,
			createdAt TIMESTAMP DEFAULT NOW(),
			updatedAt TIMESTAMP DEFAULT NOW(),
			PRIMARY KEY(id)
		)
	`

	if _, err := db.Exec(createPermissionsTable); err != nil {
		return err
	}

	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users(
			id uuid DEFAULT gen_random_uuid(),
			name VARCHAR(250) NOT NULL,
			email VARCHAR(150) NOT NULL,
			phoneNumber VARCHAR(50) NOT NULL,
			password VARCHAR(250) NOT NULL,
			deliveryMethods TEXT[] NULL,
			timeZone VARCHAR(50) DEFAULT 'Etc/UTC',
			createdAt TIMESTAMP DEFAULT NOW(),
			updatedAt TIMESTAMP NULL,
			PRIMARY KEY(id)
		)
	`

	if _, err := db.Exec(createUsersTable); err != nil {
		return err
	}

	createUserPermissionsTable := `
		CREATE TABLE IF NOT EXISTS user_permissions_permission(
			userId uuid NOT NULL,
			permissionId uuid NOT NULL,
			PRIMARY KEY(userId, permissionId),
			FOREIGN KEY(userId) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY(permissionId) REFERENCES permission(id)
		)
	`

	if _, err := db.Exec(createUserPermissionsTable); err != nil {
		return err
	}

	createApiKeysTable := `
		CREATE TABLE IF NOT EXISTS api_key(
			id uuid NOT NULL,
			name VARCHAR(250) NOT NULL,
			apiKey TEXT NOT NULL,
			expiresAt TIMESTAMP NOT NULL,
			createdBy VARCHAR(250) NOT NULL,
			createdAt TIMESTAMP DEFAULT NOW(),
			deletedBy VARCHAR(250) NULL,
			deletedAt TIMESTAMP NULL,
			PRIMARY KEY(id)
		)
	`

	if _, err := db.Exec(createApiKeysTable); err != nil {
		return err
	}

	return nil
}
