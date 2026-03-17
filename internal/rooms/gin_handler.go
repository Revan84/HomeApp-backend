package rooms

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GinHandler handles HTTP requests for rooms using Gin.
// It connects the HTTP layer to the service layer.
type GinHandler struct {
	service *Service
}

// NewGinHandler creates a new Gin handler for rooms.
func NewGinHandler(service *Service) *GinHandler {
	return &GinHandler{service: service}
}

// CreateRoom handles POST /api/v1/rooms
// It creates a new room inside a home owned by the authenticated user.
func (h *GinHandler) CreateRoom(c *gin.Context) {
	var req CreateRoomRequest

	// Bind and validate the incoming JSON payload
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// Retrieve the authenticated user ID from JWT middleware
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userID, ok := userIDRaw.(int64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id",
		})
		return
	}

	// Call the service layer to create the room
	room, err := h.service.CreateRoom(
		c.Request.Context(),
		userID,
		req.HomeID,
		req.Name,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create room",
		})
		return
	}

	// Return a clean API response
	c.JSON(http.StatusCreated, room.ToResponse())
}

// GetRoomsByHome handles GET /api/v1/homes/:homeId/rooms
// It returns all rooms belonging to a home owned by the authenticated user.
func (h *GinHandler) GetRoomsByHome(c *gin.Context) {
	// Retrieve the authenticated user ID from JWT middleware
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userID, ok := userIDRaw.(int64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id",
		})
		return
	}

	// Extract and validate homeId from route params
	homeIDParam := c.Param("homeId")
	homeID, err := strconv.ParseInt(homeIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid home id",
		})
		return
	}

	// Call the service layer to fetch rooms
	rooms, err := h.service.GetRoomsByHome(
		c.Request.Context(),
		userID,
		homeID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch rooms",
		})
		return
	}

	// Convert internal models to response DTOs
	var response []RoomResponse
	for _, room := range rooms {
		roomCopy := room
		response = append(response, roomCopy.ToResponse())
	}

	c.JSON(http.StatusOK, response)
}
