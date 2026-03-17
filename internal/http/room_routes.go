package httpapp

import (
	"github.com/gin-gonic/gin"

	"github.com/Revan84/homeapp_backend/internal/auth"
	"github.com/Revan84/homeapp_backend/internal/http/middleware"
	"github.com/Revan84/homeapp_backend/internal/rooms"
)

// registerRoomRoutes registers all routes related to rooms.
// All routes are protected by JWT authentication.
func registerRoomRoutes(
	router *gin.Engine,
	roomHandler *rooms.GinHandler,
	jwtManager *auth.JWTManager,
) {

	// API v1 group
	api := router.Group("/api/v1")

	// Apply JWT middleware to all routes in this group
	api.Use(middleware.GinAuth(jwtManager))

	// Room routes

	// Create a room
	api.POST("/rooms", roomHandler.CreateRoom)

	// Get all rooms for a specific home
	api.GET("/homes/:homeId/rooms", roomHandler.GetRoomsByHome)
}
