package httpapp

import (
	"github.com/Revan84/homeapp_backend/internal/auth"
	"github.com/Revan84/homeapp_backend/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

func registerGinAuthRoutes(
	router *gin.Engine,
	handler *auth.GinHandler,
	jwtManager *auth.JWTManager,
) {
	router.POST("/api/v1/auth/register", handler.Register)
	router.POST("/api/v1/auth/login", handler.Login)

	authGroup := router.Group("/api/v1")
	authGroup.Use(middleware.GinAuth(jwtManager))

	authGroup.GET("/me", handler.Me)

}
