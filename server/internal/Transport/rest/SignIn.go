package rest

import (
	"StillCode/server/internal/db"
	"StillCode/server/internal/models"
	_ "database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	var input models.SignUpInput

	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)", input.Email).Scan(&exists)
	if !exists {
		c.JSON(http.StatusConflict, gin.H{"error": "User isn't exist"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	_, err = db.DB.Exec("SELECT EXISTS (SELECT 1 FROM users WHERE email = $1 and password = $2)", input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password incorrect"})
		return
	}

	c.JSON(200, gin.H{"message": "Successfull"})

}
