package httpapp

import (
	"database/sql"

	"github.com/Revan84/homeapp_backend/internal/auth"
	"github.com/Revan84/homeapp_backend/internal/config"
	"github.com/gin-gonic/gin"
)

func NewGinRouter(db *sql.DB, cfg config.Config) *gin.Engine {
	router := gin.Default()

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

	authRepo := auth.NewRepository(db)
	jwtManager := &auth.JWTManager{Secret: cfg.JWTSecret}
	authService := auth.NewService(authRepo, jwtManager)
	authHandler := auth.NewGinHandler(authService)

	registerGinAuthRoutes(router, authHandler, jwtManager)

	return router
}
