package services

import (
	"StillCode/server/internal/models"
	"StillCode/server/internal/runner"
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"
)

var (
	ErrNoTestCases = errors.New("no test cases found for this task")
	ErrRunFailed   = errors.New("code execution failed")
)

// SubmissionService handles code execution and submission business logic
type SubmissionService struct {
	taskService *TaskService
}

// NewSubmissionService creates a new SubmissionService
func NewSubmissionService(taskService *TaskService) *SubmissionService {
	return &SubmissionService{
		taskService: taskService,
	}
}

// RunCode executes code with optional input and returns the result
// If TaskID is provided, the code will be wrapped with the task's function template
func (s *SubmissionService) RunCode(ctx context.Context, req *models.RunCodeRequest) (*models.RunCodeResponse, error) {
	code := req.Code

	// If TaskID is provided, wrap the code with the task template
	if req.TaskID > 0 {
		task, err := s.taskService.GetTaskForSubmission(req.TaskID)
		if err == nil && task.FunctionName != "" {
			code = s.wrapCode(req.Language, req.Code, task.FunctionName, task.Params)
		}
	}

	stdout, stderr, timeMs, status, err := runner.RunInSandbox(ctx, req.Language, code, req.Input)
	if err != nil {
		return nil, err
	}

	return &models.RunCodeResponse{
		Stdout: stdout,
		Stderr: stderr,
		TimeMs: timeMs,
		Status: status,
	}, nil
}

// RunCodeWithTask executes code for a specific task (with wrapping)
func (s *SubmissionService) RunCodeWithTask(ctx context.Context, taskID int, req *models.RunCodeRequest) (*models.RunCodeResponse, error) {
	// Get task metadata
	task, err := s.taskService.GetTaskForSubmission(taskID)
	if err != nil {
		return nil, err
	}

	// Wrap user code
	wrappedCode := s.wrapCode(req.Language, req.Code, task.FunctionName, task.Params)

	stdout, stderr, timeMs, status, err := runner.RunInSandbox(ctx, req.Language, wrappedCode, req.Input)
	if err != nil {
		return nil, err
	}

	return &models.RunCodeResponse{
		Stdout: stdout,
		Stderr: stderr,
		TimeMs: timeMs,
		Status: status,
	}, nil
}

// SubmitSolution runs code against all test cases for a task
func (s *SubmissionService) SubmitSolution(ctx context.Context, taskID int, req *models.SubmitSolutionRequest) (*models.SubmitResponse, error) {
	// Get task with ALL test cases (including hidden)
	task, err := s.taskService.GetTaskForSubmission(taskID)
	if err != nil {
		return nil, err
	}

	if len(task.TestCases) == 0 {
		return nil, ErrNoTestCases
	}

	// Wrap user code with the template
	wrappedCode := s.wrapCode(req.Language, req.Code, task.FunctionName, task.Params)

	// Run code against each test case
	var results []models.TestResult
	var totalTime int64
	passCount := 0

	for i, tc := range task.TestCases {
		stdout, stderr, durationMs, status, runErr := runner.RunInSandbox(ctx, req.Language, wrappedCode, tc.Input)
		if runErr != nil {
			return nil, runErr
		}

		totalTime += durationMs
		output := strings.TrimSpace(stdout)
		expected := strings.TrimSpace(tc.Expected)

		var resStatus string
		switch status {
		case "timeout":
			resStatus = "timeout"
		case "memory_limit":
			resStatus = "memory_limit"
		case "compile_error":
			resStatus = "compile_error"
		case "runtime_error":
			resStatus = "runtime_error"
		case "ok":
			if normalizeOutput(output) == normalizeOutput(expected) {
				resStatus = "passed"
				passCount++
			} else {
				resStatus = "wrong_answer"
			}
		default:
			resStatus = "error"
		}

		results = append(results, models.TestResult{
			TestIndex: i + 1,
			Input:     tc.Input,
			Expected:  tc.Expected,
			Output:    output + formatStderr(stderr),
			Status:    resStatus,
			TimeMs:    durationMs,
		})
	}

	return &models.SubmitResponse{
		TaskID:  taskID,
		Results: results,
		Summary: models.SubmissionSummary{
			Passed: passCount,
			Total:  len(results),
			TimeMs: totalTime,
		},
	}, nil
}

// wrapCode wraps user code with the appropriate language template
func (s *SubmissionService) wrapCode(language, code, functionName, paramsJSON string) string {
	// Parse params JSON
	var params []runner.InputParam
	if paramsJSON != "" && paramsJSON != "[]" {
		if err := json.Unmarshal([]byte(paramsJSON), &params); err != nil {
			// If parsing fails, return unwrapped code
			return code
		}
	}

	// If no function name or params, return unwrapped code
	if functionName == "" || len(params) == 0 {
		return code
	}

	wrapped := runner.WrapUserCode(language, code, functionName, params)
	// Debug: log wrapped code for Go
	if language == "go" {
		log.Println("=== WRAPPED GO CODE ===")
		log.Println(wrapped)
		log.Println("=== END WRAPPED CODE ===")
	}
	return wrapped
}

// formatStderr appends stderr if present
func formatStderr(stderr string) string {
	stderr = strings.TrimSpace(stderr)
	if stderr == "" {
		return ""
	}
	return "\n[stderr] " + stderr
}

// normalizeOutput normalizes output for comparison (handles whitespace, trailing newlines, etc.)
func normalizeOutput(s string) string {
	s = strings.TrimSpace(s)
	// Normalize JSON arrays: remove spaces after colons and commas
	s = strings.ReplaceAll(s, ", ", ",")
	s = strings.ReplaceAll(s, ": ", ":")
	return s
}
