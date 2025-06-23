package rest

import (
	"StillCode/server/internal/db"
	"StillCode/server/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProfileHandler(c *gin.Context) {

	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userID not found in context"})
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userID has wrong type"})
		return
	}

	var user models.User
	err := db.DB.QueryRow(`
        SELECT id, name, email, rating, tasks_solved 
        FROM users 
        WHERE id = $1`, userID).
		Scan(&user.Id, &user.Name, &user.Email, &user.Rating, &user.TasksSolved)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          user.Id,
		"name":        user.Name,
		"email":       user.Email,
		"rating":      user.Rating,
		"tasksSolved": user.TasksSolved,
	})
}
