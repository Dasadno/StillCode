package run

import (
	"StillCode/server/internal/db"
	"StillCode/server/internal/models"
	"StillCode/server/internal/models/runModels"
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func RunCodeHandler(c *gin.Context) {
	var req runModels.RunRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	stdout, stderr, timeMs, status, err := runInSandbox(c.Request.Context(), req.Language, req.Code, req.Input)
	if err != nil {
		// internal error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := runModels.RunResponse{
		Stdout: stdout,
		Stderr: stderr,
		TimeMs: timeMs,
		Status: status,
	}
	c.JSON(http.StatusOK, resp)
}

// Handler: POST /api/submit/:id
func SubmitSolutionHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req runModels.SubmitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	// fetch task and testcases from DB
	var t models.Task
	rows, err := db.DB.Query(`
		SELECT 
			t.id, t.title, t.description, t.difficulty, t.is_community, t.solved_percent,
			tc.input, tc.expected
		FROM tasks t
		LEFT JOIN test_cases tc ON t.id = tc.task_id
		WHERE t.id = $1
	`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var input, expected sql.NullString
		if t.ID == 0 {
			if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Difficulty, &t.IsCommunity, &t.SolvedPercent, &input, &expected); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			var dummy models.Task
			if err := rows.Scan(&dummy.ID, &dummy.Title, &dummy.Description, &dummy.Difficulty, &dummy.IsCommunity, &dummy.SolvedPercent, &input, &expected); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		if input.Valid && expected.Valid {
			t.TestCases = append(t.TestCases, models.TestCase{Input: input.String, Expected: expected.String})
		}
	}
	if t.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	// For each testcase: runInSandbox and compare output
	var results []runModels.TestResult
	var totalTime int64
	passCount := 0
	for i, tcCase := range t.TestCases {
		stdout, stderr, durationMs, status, runErr := runInSandbox(c.Request.Context(), req.Language, req.Code, tcCase.Input)
		if runErr != nil {
			// return partial error
			c.JSON(http.StatusInternalServerError, gin.H{"error": runErr.Error()})
			return
		}
		totalTime += durationMs
		outStr := stdout
		// trim trailing newlines/spaces
		outStr = strings.TrimSpace(outStr)
		expected := strings.TrimSpace(tcCase.Expected)

		resStatus := "failed"
		if status == "timeout" {
			resStatus = "timeout"
		} else if status != "ok" {
			resStatus = "error"
		} else {
			if outStr == expected {
				resStatus = "passed"
				passCount++
			} else {
				resStatus = "failed"
			}
		}

		results = append(results, runModels.TestResult{
			TestIndex: i + 1,
			Input:     tcCase.Input,
			Expected:  tcCase.Expected,
			Output:    outStr + maybeStderr(stderr),
			Status:    resStatus,
			TimeMs:    durationMs,
		})
	}

	var resp runModels.SubmitResponse
	resp.TaskID = id
	resp.Results = results
	resp.Summary.Passed = passCount
	resp.Summary.Total = len(results)
	resp.Summary.TimeMs = totalTime

	c.JSON(http.StatusOK, resp)
}
