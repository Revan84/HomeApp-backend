package httpapp

import (
	"github.com/gin-gonic/gin"

	"github.com/Revan84/homeapp_backend/internal/auth"
	"github.com/Revan84/homeapp_backend/internal/homes"
	"github.com/Revan84/homeapp_backend/internal/http/middleware"
)

// registerHomeRoutes registers all routes related to homes.
// It ensures that all routes are protected by JWT authentication.
func registerHomeRoutes(
	router *gin.Engine,
	homeHandler *homes.GinHandler,
	jwtManager *auth.JWTManager,
) {

	// Create a route group for API v1
	api := router.Group("/api/v1")

	// Apply JWT middleware to all routes in this group
	api.Use(middleware.GinAuth(jwtManager))

	// Home routes
	api.POST("/homes", homeHandler.CreateHome)
	api.GET("/homes", homeHandler.GetHomes)
}
