package rest

import (
	"StillCode/server/internal/db"
	"StillCode/server/internal/models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTaskByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	rows, err := db.DB.Query(`
        SELECT 
            t.id, t.title, t.description, t.difficulty, t.is_community, t.solved_percent,
            tc.input, tc.expected
        FROM tasks t
        LEFT JOIN test_cases tc ON t.id = tc.task_id
        WHERE t.id = $1
    `, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database query error"})
		return
	}
	defer rows.Close()

	var t models.Task
	for rows.Next() {
		var (
			input, expected sql.NullString
		)

		if t.ID == 0 {
			err := rows.Scan(
				&t.ID, &t.Title, &t.Description, &t.Difficulty, &t.IsCommunity, &t.SolvedPercent,
				&input, &expected,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error reading task"})
				return
			}
		} else {

			var dummy models.Task
			err := rows.Scan(
				&dummy.ID, &dummy.Title, &dummy.Description, &dummy.Difficulty,
				&dummy.IsCommunity, &dummy.SolvedPercent, &input, &expected,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error reading test cases"})
				return
			}
		}

		if input.Valid && expected.Valid {
			t.TestCases = append(t.TestCases, models.TestCase{
				Input:    input.String,
				Expected: expected.String,
			})
		}
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "iteration error"})
		return
	}

	if t.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	c.JSON(http.StatusOK, t)
}
