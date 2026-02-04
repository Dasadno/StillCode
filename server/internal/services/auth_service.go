package services

import (
	"StillCode/server/internal/auth"
	"StillCode/server/internal/db"
	"StillCode/server/internal/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserExists         = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrHashFailed         = errors.New("failed to hash password")
	ErrDBOperation        = errors.New("database operation failed")
	ErrTokenGeneration    = errors.New("failed to generate token")
)

// AuthService handles authentication business logic
type AuthService struct{}

// NewAuthService creates a new AuthService
func NewAuthService() *AuthService {
	return &AuthService{}
}

// SignUp registers a new user
func (s *AuthService) SignUp(req *models.SignUpRequest) error {
	// Check if user already exists
	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", req.Email).Scan(&exists)
	if err != nil {
		return ErrDBOperation
	}
	if exists {
		return ErrUserExists
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return ErrHashFailed
	}

	// Insert user
	_, err = db.DB.Exec(
		`INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`,
		req.Name, req.Email, string(hashedPassword),
	)
	if err != nil {
		return ErrDBOperation
	}

	return nil
}

// SignIn authenticates a user and returns a JWT token
func (s *AuthService) SignIn(req *models.SignInRequest) (string, error) {
	var userID int
	var hashedPassword string

	err := db.DB.QueryRow(
		`SELECT id, password FROM users WHERE email = $1`,
		req.Email,
	).Scan(&userID, &hashedPassword)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		return "", ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := auth.GenerateJWT(userID)
	if err != nil {
		return "", ErrTokenGeneration
	}

	return token, nil
}
