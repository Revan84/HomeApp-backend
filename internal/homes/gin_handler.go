package homes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GinHandler handles HTTP requests for homes using Gin.
// It connects the HTTP layer to the service layer.
type GinHandler struct {
	service *Service
}

// NewGinHandler creates a new Gin handler for homes.
func NewGinHandler(service *Service) *GinHandler {
	return &GinHandler{service: service}
}

// CreateHome handles POST /api/v1/homes
// It creates a new home for the authenticated user.
func (h *GinHandler) CreateHome(c *gin.Context) {

	var req CreateHomeRequest

	// Bind and validate JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// Retrieve userID from JWT middleware
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

	// Call service layer
	home, err := h.service.CreateHome(
		c.Request.Context(),
		userID,
		req.Name,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create home",
		})
		return
	}

	// Return response (without exposing internal fields)
	c.JSON(http.StatusCreated, home.ToResponse())
}

// GetHomes handles GET /api/v1/homes
// It retrieves all homes belonging to the authenticated user.
func (h *GinHandler) GetHomes(c *gin.Context) {

	// Retrieve userID from JWT middleware
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

	// Call service layer
	homes, err := h.service.GetHomesByUser(
		c.Request.Context(),
		userID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch homes",
		})
		return
	}

	// Convert to response DTO
	var response []HomeResponse
	for _, h := range homes {
		response = append(response, h.ToResponse())
	}

	c.JSON(http.StatusOK, response)
}
