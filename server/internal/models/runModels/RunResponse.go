package runModels

type RunResponse struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	TimeMs int64  `json:"time_ms"`
	Status string `json:"status"`
}
