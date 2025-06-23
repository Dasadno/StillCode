package rest

import (
	"StillCode/server/internal/auth"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func SetupHandlers(r *gin.Engine) {

	r.POST("/signup", RegisterHandler)
	r.POST("/signin", SignInHandler)

	protected := r.Group("/api")
	protected.Use(auth.AuthRequired())
	protected.GET("/profile", ProfileHandler)
}
