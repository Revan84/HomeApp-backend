package homes

import "context"

// Service contains business logic related to homes.
// It acts as a layer between handlers (HTTP) and repository (DB).
type Service struct {
	repo *Repository
}

// NewService creates a new Home service instance.
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// CreateHome handles the business logic for creating a home.
// It ensures the home is linked to the correct user.
func (s *Service) CreateHome(
	ctx context.Context,
	userID int64,
	name string,
) (*Home, error) {

	// Future-proof: here we can add business rules
	// Example:
	// - limit number of homes per user
	// - sanitize name
	// - check duplicates

	return s.repo.Create(ctx, userID, name)
}

// GetHomesByUser retrieves all homes for a given user.
// It ensures data isolation between users.
func (s *Service) GetHomesByUser(
	ctx context.Context,
	userID int64,
) ([]Home, error) {

	return s.repo.FindByUserID(ctx, userID)
}
