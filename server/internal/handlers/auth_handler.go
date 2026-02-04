package handlers

import (
	"StillCode/server/internal/models"
	"StillCode/server/internal/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication HTTP requests
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// SignUp handles POST /api/auth/signup
func (h *AuthHandler) SignUp(c *gin.Context) {
	var req models.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input"})
		return
	}

	err := h.authService.SignUp(&req)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrUserExists):
			c.JSON(http.StatusConflict, models.ErrorResponse{Error: "User already exists"})
		case errors.Is(err, services.ErrHashFailed):
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Password hash failed"})
		default:
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Registration failed"})
		}
		return
	}

	c.JSON(http.StatusCreated, models.MessageResponse{Message: "Registration successful"})
}

// SignIn handles POST /api/auth/signin
func (h *AuthHandler) SignIn(c *gin.Context) {
	var req models.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input"})
		return
	}

	token, err := h.authService.SignIn(&req)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrInvalidCredentials):
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid email or password"})
		case errors.Is(err, services.ErrTokenGeneration):
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Token generation failed"})
		default:
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Sign in failed"})
		}
		return
	}

	c.JSON(http.StatusOK, models.AuthResponse{Token: token})
}
