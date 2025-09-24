package models

type Task struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Difficulty    string  `json:"difficulty"`
	IsCommunity   bool    `json:"isCommunity"`
	SolvedPercent float64 `json:"solvedPercent"`
}
