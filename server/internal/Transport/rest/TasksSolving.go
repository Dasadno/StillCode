package rest

import (
	"StillCode/server/internal/db"
	"StillCode/server/internal/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTaskByIDHandler(c *gin.Context) {
	id := c.Param("id")

	var t models.Task
	var test models.TestCase
	err := db.DB.QueryRow(`
        SELECT id, title, description, difficulty, is_community, solved_percent
        FROM tasks
        WHERE id = $1
    `, id).Scan(&t.ID, &t.Title, &t.Description, &t.Difficulty, &t.IsCommunity, &t.SolvedPercent)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		}
		return
	}

	rows, err := db.DB.Query(`
        SELECT input, expected
        FROM test_cases
        WHERE task_id = $1
    `, id)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var tc models.TestCase
			rows.Scan(&tc.Input, &tc.Expected)
			test.TestCase = append(tc.TestCase, tc)
		}
	}

	c.JSON(http.StatusOK, t)
}
