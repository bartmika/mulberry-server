package repositories

import (
	"context"
	"database/sql"
	// "encoding/json"
	"time"

	"github.com/bartmika/mulberry-server/pkg/models"
)

// UserRepo implements models.UserRepository
type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(ctx context.Context, uuid string, name string, email string, passwordHash string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "INSERT INTO users (uuid, name, email, password_hash) VALUES ($1, $2, $3, $4)"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		uuid,
		name,
		email,
		passwordHash,
	)
	return err
}

func (r *UserRepo) FindByUuid(ctx context.Context, uuid string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.User)

	query := "SELECT uuid, name, email, password_hash FROM users WHERE uuid = $1"
	err := r.db.QueryRowContext(ctx, query, uuid).Scan(
		&m.Uuid,
		&m.Name,
		&m.Email,
		&m.PasswordHash,
	)
	if err != nil {
		// CASE 1 OF 2: Cannot find record with that uuid.
		if err == sql.ErrNoRows {
			return nil, nil
		} else { // CASE 2 OF 2: All other errors.
			return nil, err
		}
	}
	return m, nil
}

func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.User)

	query := "SELECT uuid, name, email, password_hash FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&m.Uuid,
		&m.Name,
		&m.Email,
		&m.PasswordHash,
	)
	if err != nil {
		// CASE 1 OF 2: Cannot find record with that email.
		if err == sql.ErrNoRows {
			return nil, nil
		} else { // CASE 2 OF 2: All other errors.
			return nil, err
		}
	}
	return m, nil
}

func (r *UserRepo) Save(ctx context.Context, m *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "UPDATE users SET name = $1, email = $2, password_hash = $3 WHERE uuid = $4"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		m.Name,
		m.Email,
		m.PasswordHash,
		m.Uuid,
	)
	return err
}
