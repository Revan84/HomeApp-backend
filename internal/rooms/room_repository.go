package rooms

import (
	"context"
	"database/sql"
)

// Repository handles database operations for rooms.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new Room repository instance.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// Create inserts a new room into the database.
// It ensures the home belongs to the user before creating the room.
func (r *Repository) Create(
	ctx context.Context,
	userID int64,
	homeID int64,
	name string,
) (*Room, error) {

	var room Room

	query := `
		INSERT INTO rooms (home_id, name)
		SELECT h.id, $3
		FROM homes h
		WHERE h.id = $1 AND h.user_id = $2
		RETURNING id, home_id, name, created_at
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		homeID,
		userID,
		name,
	).Scan(
		&room.ID,
		&room.HomeID,
		&room.Name,
		&room.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &room, nil
}

// FindByHomeID retrieves all rooms for a given home,
// but only if the home belongs to the user.
func (r *Repository) FindByHomeID(
	ctx context.Context,
	userID int64,
	homeID int64,
) ([]Room, error) {

	query := `
		SELECT r.id, r.home_id, r.name, r.created_at
		FROM rooms r
		INNER JOIN homes h ON r.home_id = h.id
		WHERE r.home_id = $1 AND h.user_id = $2
		ORDER BY r.created_at ASC
	`

	rows, err := r.db.QueryContext(ctx, query, homeID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []Room

	for rows.Next() {
		var r Room

		err := rows.Scan(
			&r.ID,
			&r.HomeID,
			&r.Name,
			&r.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		rooms = append(rooms, r)
	}

	// Check for iteration errors
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}
