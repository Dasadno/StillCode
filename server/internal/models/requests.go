package models

// SignUpRequest represents the signup request payload
type SignUpRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// SignInRequest represents the signin request payload
type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RunCodeRequest represents the request to run code
type RunCodeRequest struct {
	Language string `json:"language" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Input    string `json:"input"`
}

// SubmitSolutionRequest represents the request to submit a solution
type SubmitSolutionRequest struct {
	Language string `json:"language" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

// TasksQueryParams represents query parameters for listing tasks
type TasksQueryParams struct {
	Search     string `form:"search"`
	Difficulty string `form:"difficulty"`
	Community  string `form:"community"`
	Page       int    `form:"page,default=1"`
	PageSize   int    `form:"pageSize,default=20"`
}
