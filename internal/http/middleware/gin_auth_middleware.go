package middleware

import (
	"strings"

	"github.com/Revan84/homeapp_backend/internal/auth"
	"github.com/gin-gonic/gin"
)

func GinAuth(jwtManager *auth.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "missing authorization header"})
			return
		}

		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid authorization header"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, bearerPrefix)

		claims, err := jwtManager.Parse(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
