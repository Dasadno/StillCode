package runModels

type SubmitResponse struct {
	TaskID  int          `json:"task_id"`
	Results []TestResult `json:"results"`

	Summary struct {
		Passed int   `json:"passed"`
		Total  int   `json:"total"`
		TimeMs int64 `json:"time_ms"`
	} `json:"summary"`
}
