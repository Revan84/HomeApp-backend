package httpapp

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/Revan84/homeapp_backend/internal/auth"
	"github.com/Revan84/homeapp_backend/internal/config"
	"github.com/Revan84/homeapp_backend/internal/homes"
	"github.com/Revan84/homeapp_backend/internal/rooms"
)

// NewGinRouter initializes the Gin router and registers all routes.
func NewGinRouter(db *sql.DB, cfg config.Config) *gin.Engine {

	router := gin.Default()

	// =========================
	// HEALTH CHECKS
	// =========================
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.GET("/health/db", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.JSON(503, gin.H{"status": "db_down"})
			return
		}
		c.JSON(200, gin.H{"status": "db_up"})
	})

	// =========================
	// AUTH MODULE
	// =========================
	authRepo := auth.NewRepository(db)
	jwtManager := &auth.JWTManager{Secret: cfg.JWTSecret}
	authService := auth.NewService(authRepo, jwtManager)
	authHandler := auth.NewGinHandler(authService)

	registerGinAuthRoutes(router, authHandler, jwtManager)

	// =========================
	// HOMES MODULE
	// =========================
	homeRepo := homes.NewRepository(db)
	homeService := homes.NewService(homeRepo)
	homeHandler := homes.NewGinHandler(homeService)

	registerHomeRoutes(router, homeHandler, jwtManager)

	// =========================
	// ROOMS MODULE
	// =========================
	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewGinHandler(roomService)

	registerRoomRoutes(router, roomHandler, jwtManager)

	return router
}
