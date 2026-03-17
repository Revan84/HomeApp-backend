package rooms

import "context"

// Service contains business logic related to rooms.
// It acts as a bridge between the HTTP layer (handlers) and the repository (database).
type Service struct {
	repo *Repository
}

// NewService creates a new Room service instance.
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// CreateRoom handles the business logic for creating a room.
// It ensures the room is created only if the home belongs to the user.
func (s *Service) CreateRoom(
	ctx context.Context,
	userID int64,
	homeID int64,
	name string,
) (*Room, error) {

	// Future business rules can be added here:
	// - limit number of rooms per home
	// - prevent duplicate room names within the same home
	// - sanitize input (trim, normalize, etc.)

	return s.repo.Create(ctx, userID, homeID, name)
}

// GetRoomsByHome retrieves all rooms for a given home,
// ensuring the home belongs to the user.
func (s *Service) GetRoomsByHome(
	ctx context.Context,
	userID int64,
	homeID int64,
) ([]Room, error) {

	return s.repo.FindByHomeID(ctx, userID, homeID)
}
