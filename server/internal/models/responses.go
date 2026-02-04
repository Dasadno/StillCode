package models

// AuthResponse represents the response after successful authentication
type AuthResponse struct {
	Token string `json:"token"`
}

// MessageResponse represents a simple message response
type MessageResponse struct {
	Message string `json:"message"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error  string `json:"error"`
	Detail string `json:"detail,omitempty"`
}

// ProfileResponse represents the user profile response
type ProfileResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Rating      int    `json:"rating"`
	TasksSolved int    `json:"tasksSolved"`
}

// RunCodeResponse represents the response from running code
type RunCodeResponse struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	TimeMs int64  `json:"time_ms"`
	Status string `json:"status"`
}

// SubmitResponse represents the response from submitting a solution
type SubmitResponse struct {
	TaskID  int               `json:"task_id"`
	Results []TestResult      `json:"results"`
	Summary SubmissionSummary `json:"summary"`
}
