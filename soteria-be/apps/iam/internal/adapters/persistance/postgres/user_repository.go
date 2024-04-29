package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/domain"
	"github.com/cmclaughlin24/soteria-be/pkg/iam"
)

type userPermissionEntity struct {
	Resource sql.NullString
	Action   sql.NullString
}

type PgxUserRepository struct {
	db *sql.DB
}

func NewUserRepository(connection *sql.DB) *PgxUserRepository {
	return &PgxUserRepository{
		db: connection,
	}
}

func (r *PgxUserRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	return r.find(ctx, struct {
		id    string
		email string
	}{})
}

func (r *PgxUserRepository) FindOne(ctx context.Context, id string) (*domain.User, error) {
	users, err := r.find(ctx, struct {
		id    string
		email string
	}{id: id})

	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return &users[0], nil
}

func (r *PgxUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	users, err := r.find(ctx, struct {
		id    string
		email string
	}{email: email})

	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return &users[0], nil
}

func (r *PgxUserRepository) Create(ctx context.Context, u domain.User) (*domain.User, error) {
	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO users(
			name,
			email,
			phoneNumber,
			password,
			deliveryMethods,
			timeZone,
			createdAt
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(
		u.Name,
		u.Email,
		u.PhoneNumber,
		u.Password,
		u.DeliveryMethods,
		u.TimeZone,
		time.Now(),
	)

	var userId string

	if err := row.Scan(&userId); err != nil {
		return nil, err
	}

	permissions, err := r.findPermissions(ctx, u.Permissions)

	if err != nil {
		return nil, err
	}

	if err = r.createUserPermissions(tx, userId, permissions); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	u.Id = userId

	return &u, nil
}

func (r *PgxUserRepository) Update(ctx context.Context, u domain.User) (*domain.User, error) {
	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		UPDATE users
		SET name = $2,
			email = $3,
			phoneNumber = $4,
			deliveryMethods = $5,
			timeZone = $6,
			updatedAt = $7
		WHERE id = $1
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		u.Id,
		u.Name,
		u.Email,
		u.PhoneNumber,
		u.DeliveryMethods,
		u.TimeZone,
		time.Now(),
	)

	if err != nil {
		return nil, err
	}

	_, err = tx.Exec("DELETE FROM user_permissions_permission WHERE userId = $1", u.Id)

	if err != nil {
		return nil, err
	}

	permissions, err := r.findPermissions(ctx, u.Permissions)

	if err != nil {
		return nil, err
	}

	if err = r.createUserPermissions(tx, u.Id, permissions); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *PgxUserRepository) Remove(ctx context.Context, id string) error {
	stmt, err := r.db.PrepareContext(ctx, "DELETE FROM users WHERE id = $1")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func (r *PgxUserRepository) find(ctx context.Context, options struct {
	id    string
	email string
}) ([]domain.User, error) {
	query := `
		SELECT users.id,
			users.name,
			users.email,
			users.phoneNumber,
			users.deliveryMethods,
			users.password,
			users.timeZone,
			permission.resource,
			permission.action
		FROM users
			LEFT JOIN (
				user_permissions_permission
				LEFT JOIN permission ON permissionid = permission.id
			) ON users.id = userid
	`
	args := make([]any, 0)

	if options.id != "" {
		query += "WHERE users.id = $1"
		args = append(args, options.id)
	} else if options.email != "" {
		query += "WHERE users.email = $1"
		args = append(args, options.email)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	entities := make(map[string]domain.User)

	for rows.Next() {
		var u domain.User
		var p userPermissionEntity
		var deliveryMethods string

		err = rows.Scan(
			&u.Id,
			&u.Name,
			&u.Email,
			&u.PhoneNumber,
			&deliveryMethods,
			&u.Password,
			&u.TimeZone,
			&p.Resource,
			&p.Action,
		)

		if err != nil {
			return nil, err
		}

		if existing, ok := entities[u.Id]; ok {
			u = existing
		}

		if p.Resource.Valid && p.Action.Valid {
			u.AddPermission(iam.UserPermission{
				Resource: p.Resource.String,
				Action:   p.Action.String,
			})
		}

		u.DeliveryMethods = parsePostgresTextArray(deliveryMethods)
		entities[u.Id] = u
	}

	users := make([]domain.User, 0, len(entities))

	for _, u := range entities {
		users = append(users, u)
	}

	return users, nil
}

func (r *PgxUserRepository) findPermissions(ctx context.Context, userPermissions []iam.UserPermission) ([]string, error) {
	query := "SELECT id FROM permission WHERE resource = $1 AND action = $2"
	permissions := make([]string, 0)

	for _, permission := range userPermissions {
		var id string
		row := r.db.QueryRowContext(ctx, query, permission.Resource, permission.Action)

		if err := row.Scan(&id); err != nil {
			return nil, fmt.Errorf(
				"error=\"%s\" occurred when looking up permission resource=%s action=%s",
				err.Error(),
				permission.Resource,
				permission.Action,
			)
		}

		permissions = append(permissions, id)
	}

	return permissions, nil
}

func (r *PgxUserRepository) createUserPermissions(tx *sql.Tx, userId string, permissionIds []string) error {
	stmt, err := tx.Prepare(`
		INSERT INTO user_permissions_permission(
			userId,
			permissionId
		) VALUES ($1, $2)
	`)

	if err != nil {
		return err
	}

	for _, permissionId := range permissionIds {
		_, err := stmt.Exec(userId, permissionId)

		if err != nil {
			return err
		}
	}

	return nil
}
