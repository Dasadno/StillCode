package models

// User represents a user in the system
type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"-"` // Never send password in JSON
	Rating      int    `json:"rating"`
	TasksSolved int    `json:"tasks_solved"`
}
