package auth

import (
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(
	ctx context.Context,
	email string,
	passwordHash string,
) (int64, error) {

	var id int64

	err := r.db.QueryRowContext(
		ctx,
		`INSERT INTO users (email, password_hash)
		 VALUES ($1,$2)
		 RETURNING id`,
		email,
		passwordHash,
	).Scan(&id)

	return id, err
}

func (r *Repository) FindByEmail(
	ctx context.Context,
	email string,
) (int64, string, error) {

	var id int64
	var hash string

	err := r.db.QueryRowContext(
		ctx,
		`SELECT id,password_hash FROM users WHERE email=$1`,
		email,
	).Scan(&id, &hash)

	return id, hash, err
}
