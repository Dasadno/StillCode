package models

type TestCase struct {
	TestCase string `json:"testCase"`
	Input    string `json:"input"`
	Expected string `json:"expected"`
}
