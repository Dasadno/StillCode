package services

import (
	"StillCode/server/internal/db"
	"StillCode/server/internal/models"
	"database/sql"
	"errors"
	"strconv"
	"strings"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

// TaskService handles task-related business logic
type TaskService struct{}

// NewTaskService creates a new TaskService
func NewTaskService() *TaskService {
	return &TaskService{}
}

// GetTasks retrieves a list of tasks with optional filtering and pagination
func (s *TaskService) GetTasks(params *models.TasksQueryParams) ([]models.Task, error) {
	whereClauses := []string{}
	args := []interface{}{}
	idx := 1

	if params.Search != "" {
		whereClauses = append(whereClauses, "title ILIKE '%' || $"+strconv.Itoa(idx)+" || '%'")
		args = append(args, params.Search)
		idx++
	}
	if params.Difficulty != "" {
		whereClauses = append(whereClauses, "difficulty = $"+strconv.Itoa(idx))
		args = append(args, params.Difficulty)
		idx++
	}
	if params.Community != "" {
		whereClauses = append(whereClauses, "is_community = $"+strconv.Itoa(idx))
		b, _ := strconv.ParseBool(params.Community)
		args = append(args, b)
		idx++
	}

	query := "SELECT id, title, difficulty, is_community, solved_percent FROM tasks"
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	offset := (params.Page - 1) * params.PageSize
	query += " ORDER BY id LIMIT $" + strconv.Itoa(idx) + " OFFSET $" + strconv.Itoa(idx+1)
	args = append(args, params.PageSize, offset)

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, ErrDBOperation
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Difficulty, &t.IsCommunity, &t.SolvedPercent); err != nil {
			continue
		}
		tasks = append(tasks, t)
	}

	if tasks == nil {
		tasks = []models.Task{}
	}

	return tasks, nil
}

// GetTaskByID retrieves a task by its ID with test cases and starter code
func (s *TaskService) GetTaskByID(taskID int) (*models.Task, error) {
	// First get the task details
	var task models.Task
	var starterGo, starterJS, starterPython, starterCpp, starterJava sql.NullString

	err := db.DB.QueryRow(`
		SELECT id, title, description, difficulty, is_community, solved_percent,
		       function_name, params,
		       starter_code_go, starter_code_js, starter_code_python, starter_code_cpp, starter_code_java
		FROM tasks WHERE id = $1
	`, taskID).Scan(
		&task.ID, &task.Title, &task.Description, &task.Difficulty,
		&task.IsCommunity, &task.SolvedPercent,
		&task.FunctionName, &task.Params,
		&starterGo, &starterJS, &starterPython, &starterCpp, &starterJava,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrTaskNotFound
		}
		return nil, ErrDBOperation
	}

	// Build starter code map
	task.StarterCode = make(map[string]string)
	if starterGo.Valid {
		task.StarterCode["go"] = starterGo.String
	}
	if starterJS.Valid {
		task.StarterCode["javascript"] = starterJS.String
	}
	if starterPython.Valid {
		task.StarterCode["python"] = starterPython.String
	}
	if starterCpp.Valid {
		task.StarterCode["cpp"] = starterCpp.String
	}
	if starterJava.Valid {
		task.StarterCode["java"] = starterJava.String
	}

	// Get test cases (only non-hidden ones for the response)
	rows, err := db.DB.Query(`
		SELECT input, expected FROM test_cases
		WHERE task_id = $1 AND is_hidden = FALSE
		ORDER BY id
	`, taskID)
	if err != nil {
		return nil, ErrDBOperation
	}
	defer rows.Close()

	for rows.Next() {
		var tc models.TestCase
		if err := rows.Scan(&tc.Input, &tc.Expected); err != nil {
			continue
		}
		task.TestCases = append(task.TestCases, tc)
	}

	return &task, nil
}

// GetTaskForSubmission retrieves task with ALL test cases (including hidden) for submission
func (s *TaskService) GetTaskForSubmission(taskID int) (*models.Task, error) {
	var task models.Task

	err := db.DB.QueryRow(`
		SELECT id, title, function_name, params FROM tasks WHERE id = $1
	`, taskID).Scan(&task.ID, &task.Title, &task.FunctionName, &task.Params)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrTaskNotFound
		}
		return nil, ErrDBOperation
	}

	// Get ALL test cases including hidden
	rows, err := db.DB.Query(`
		SELECT input, expected FROM test_cases
		WHERE task_id = $1
		ORDER BY id
	`, taskID)
	if err != nil {
		return nil, ErrDBOperation
	}
	defer rows.Close()

	for rows.Next() {
		var tc models.TestCase
		if err := rows.Scan(&tc.Input, &tc.Expected); err != nil {
			continue
		}
		task.TestCases = append(task.TestCases, tc)
	}

	return &task, nil
}
