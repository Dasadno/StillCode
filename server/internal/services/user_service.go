package services

import (
	"StillCode/server/internal/db"
	"StillCode/server/internal/models"
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

// UserService handles user-related business logic
type UserService struct{}

// NewUserService creates a new UserService
func NewUserService() *UserService {
	return &UserService{}
}

// GetProfile retrieves a user's profile by their ID
func (s *UserService) GetProfile(userID int) (*models.ProfileResponse, error) {
	var profile models.ProfileResponse

	err := db.DB.QueryRow(`
		SELECT id, name, email, rating, tasks_solved
		FROM users
		WHERE id = $1
	`, userID).Scan(&profile.ID, &profile.Name, &profile.Email, &profile.Rating, &profile.TasksSolved)

	if err != nil {
		return nil, ErrUserNotFound
	}

	return &profile, nil
}

// GetUserStats retrieves statistics for a user
func (s *UserService) GetUserStats(userID int) (map[string]interface{}, error) {
	var totalSubmissions, acceptedSubmissions int

	err := db.DB.QueryRow(`
		SELECT
			COUNT(*) as total,
			COUNT(*) FILTER (WHERE status = 'accepted') as accepted
		FROM submissions
		WHERE user_id = $1
	`, userID).Scan(&totalSubmissions, &acceptedSubmissions)

	if err != nil {
		return nil, ErrDBOperation
	}

	return map[string]interface{}{
		"total_submissions":    totalSubmissions,
		"accepted_submissions": acceptedSubmissions,
	}, nil
}
