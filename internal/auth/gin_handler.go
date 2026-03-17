package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	service *Service
}

func NewGinHandler(service *Service) *GinHandler {
	return &GinHandler{service: service}
}

func (h *GinHandler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	token, err := h.service.Register(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to register",
		})
		return
	}

	c.JSON(http.StatusOK, AuthResponse{
		Token: token,
	})
}

func (h *GinHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	token, err := h.service.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	c.JSON(http.StatusOK, AuthResponse{
		Token: token,
	})
}

func (h *GinHandler) Me(c *gin.Context) {
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID, ok := userIDRaw.(int64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user id"})
		return
	}
	user, err := h.service.Me(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch user",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
