package models

// TestCase represents a test case for a task
type TestCase struct {
	ID       int    `json:"id,omitempty"`
	TaskID   int    `json:"task_id,omitempty"`
	Input    string `json:"input"`
	Expected string `json:"expected"`
	IsHidden bool   `json:"is_hidden,omitempty"`
}
