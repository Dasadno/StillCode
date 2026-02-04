package handlers

import (
	"StillCode/server/internal/models"
	"StillCode/server/internal/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetProfile handles GET /api/profile
func (h *UserHandler) GetProfile(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "User ID not found in context"})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "User ID has wrong type"})
		return
	}

	profile, err := h.userService.GetProfile(userID)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to fetch profile"})
		return
	}

	c.JSON(http.StatusOK, profile)
}
