package homes

import (
	"context"
	"database/sql"
)

// Repository handles database operations for homes.
// It is responsible for interacting with the PostgreSQL database.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new Home repository instance.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// Create inserts a new home into the database and returns the created Home.
// It links the home to a specific user via userID.
func (r *Repository) Create(
	ctx context.Context,
	userID int64,
	name string,
) (*Home, error) {

	var home Home

	query := `
		INSERT INTO homes (user_id, name)
		VALUES ($1, $2)
		RETURNING id, user_id, name, created_at
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		userID,
		name,
	).Scan(
		&home.ID,
		&home.UserID,
		&home.Name,
		&home.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &home, nil
}

// FindByUserID retrieves all homes belonging to a specific user.
// This ensures that users can only access their own homes.
func (r *Repository) FindByUserID(
	ctx context.Context,
	userID int64,
) ([]Home, error) {

	query := `
		SELECT id, user_id, name, created_at
		FROM homes
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var homes []Home

	for rows.Next() {
		var h Home

		err := rows.Scan(
			&h.ID,
			&h.UserID,
			&h.Name,
			&h.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		homes = append(homes, h)
	}

	// Check for iteration errors
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return homes, nil
}
