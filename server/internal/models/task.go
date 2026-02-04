package models

// Task represents a coding task/problem
type Task struct {
	ID            int               `json:"id"`
	Title         string            `json:"title"`
	Description   string            `json:"description,omitempty"`
	Difficulty    string            `json:"difficulty"`
	IsCommunity   bool              `json:"isCommunity"`
	SolvedPercent float64           `json:"solvedPercent"`
	FunctionName  string            `json:"functionName,omitempty"`
	Params        string            `json:"params,omitempty"` // JSON: [{"name":"nums","type":"[]int"},{"name":"target","type":"int"}]
	StarterCode   map[string]string `json:"starterCode,omitempty"`
	TestCases     []TestCase        `json:"test_cases,omitempty"`
}

// FunctionParam represents a function parameter
type FunctionParam struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
