package rest

import (
	"StillCode/server/internal/db"
	"StillCode/server/internal/models"
	_ "database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c *gin.Context) {
	var input models.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var exists bool
	db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM public.users WHERE email = $1)", input.Email).Scan(&exists)
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hash failed"})
		return
	}

	_, err = db.DB.Exec(`INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`,
		input.Name, input.Email, string(hashedPassword))
	if err != nil {
		log.Printf("RegisterHandler: insert error: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "DB insert failed",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}
