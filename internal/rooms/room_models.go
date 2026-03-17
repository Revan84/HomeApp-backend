package rooms

import "time"

// Room represents a room inside a home.
// A room always belongs to a specific home.
type Room struct {
	ID        int64     `json:"id"`         // Unique identifier of the room
	HomeID    int64     `json:"home_id"`    // Parent home identifier
	Name      string    `json:"name"`       // Room name (e.g. "Living room", "Kitchen")
	CreatedAt time.Time `json:"created_at"` // Creation timestamp
}

// CreateRoomRequest represents the payload required to create a room.
// Validation is handled by Gin using struct tags.
type CreateRoomRequest struct {
	HomeID int64  `json:"home_id" binding:"required"`           // Target home identifier
	Name   string `json:"name" binding:"required,min=2,max=80"` // Room name must be between 2 and 80 characters
}

// RoomResponse represents the room data exposed to the client.
// It avoids leaking internal fields unnecessarily.
type RoomResponse struct {
	ID        int64     `json:"id"`
	HomeID    int64     `json:"home_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// ToResponse converts a Room model into a RoomResponse.
func (r *Room) ToResponse() RoomResponse {
	return RoomResponse{
		ID:        r.ID,
		HomeID:    r.HomeID,
		Name:      r.Name,
		CreatedAt: r.CreatedAt,
	}
}
