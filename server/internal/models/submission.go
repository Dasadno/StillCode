package models

import "time"

// Submission represents a code submission for a task
type Submission struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	TaskID      int       `json:"task_id"`
	Language    string    `json:"language"`
	Code        string    `json:"code"`
	Status      string    `json:"status"`
	RuntimeMs   int       `json:"runtime_ms"`
	MemoryKb    int       `json:"memory_kb"`
	PassedTests int       `json:"passed_tests"`
	TotalTests  int       `json:"total_tests"`
	CreatedAt   time.Time `json:"created_at"`
}

// TestResult represents the result of running a single test case
type TestResult struct {
	TestIndex int    `json:"test_index"`
	Input     string `json:"input"`
	Expected  string `json:"expected"`
	Output    string `json:"output"`
	Status    string `json:"status"`
	TimeMs    int64  `json:"time_ms"`
}

// SubmissionSummary provides a summary of submission results
type SubmissionSummary struct {
	Passed int   `json:"passed"`
	Total  int   `json:"total"`
	TimeMs int64 `json:"time_ms"`
}
