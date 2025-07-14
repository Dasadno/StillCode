package rest

import (
	"StillCode/server/internal/db"
	"StillCode/server/internal/models"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetTasksHandler(c *gin.Context) {
	search := c.Query("search")
	difficulty := c.Query("difficulty")
	community := c.Query("community")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	offset := (page - 1) * size

	whereClauses := []string{}
	args := []interface{}{}
	idx := 1

	if search != "" {
		whereClauses = append(whereClauses, "title ILIKE '%' || $"+strconv.Itoa(idx)+" || '%'")
		args = append(args, search)
		idx++
	}
	if difficulty != "" {
		whereClauses = append(whereClauses, "difficulty = $"+strconv.Itoa(idx))
		args = append(args, difficulty)
		idx++
	}
	if community != "" {
		whereClauses = append(whereClauses, "is_community = $"+strconv.Itoa(idx))
		b, _ := strconv.ParseBool(community)
		args = append(args, b)
		idx++
	}

	sql := "SELECT id, title, difficulty, is_community, solved_percent FROM tasks"
	if len(whereClauses) > 0 {
		sql += " WHERE " + strings.Join(whereClauses, " AND ")
	}
	sql += " ORDER BY id LIMIT $" + strconv.Itoa(idx) + " OFFSET $" + strconv.Itoa(idx+1)
	args = append(args, size, offset)

	log.Printf("GetTasksHandler: search=%q difficulty=%q community=%q page=%d size=%d",
		search, difficulty, community, page, size)

	rows, err := db.DB.Query(sql, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
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
	c.JSON(http.StatusOK, tasks)
}
