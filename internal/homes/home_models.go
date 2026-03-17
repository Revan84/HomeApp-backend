package homes

import "time"

// Home represents a housing entity owned by a user.
// It is the main container for rooms and equipments.
type Home struct {
	ID        int64     `json:"id"`         // Unique identifier of the home
	UserID    int64     `json:"user_id"`    // Owner of the home (foreign key to users table)
	Name      string    `json:"name"`       // Name of the home (e.g. "Apartment", "House")
	CreatedAt time.Time `json:"created_at"` // Timestamp of creation
}

// CreateHomeRequest represents the payload required to create a new home.
// Validation is handled by Gin using struct tags.
type CreateHomeRequest struct {
	Name string `json:"name" binding:"required,min=2,max=100"` // Home name must be between 2 and 100 characters
}

// HomeResponse represents the data returned to the client.
// It allows you to control what fields are exposed (good practice).
type HomeResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// ToResponse converts a Home model into a HomeResponse.
// This avoids leaking internal fields like UserID.
func (h *Home) ToResponse() HomeResponse {
	return HomeResponse{
		ID:        h.ID,
		Name:      h.Name,
		CreatedAt: h.CreatedAt,
	}
}
