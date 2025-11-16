package runModels

type TestResult struct {
	TestIndex int    `json:"test_index"`
	Input     string `json:"input"`
	Expected  string `json:"expected"`
	Output    string `json:"output"`
	Status    string `json:"status"`
	TimeMs    int64  `json:"time_ms"`
}
