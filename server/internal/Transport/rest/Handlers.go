package rest

import (
	"StillCode/server/internal/Transport/rest/run"
	"StillCode/server/internal/auth"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func SetupHandlers(r *gin.Engine) {

	r.POST("/api/run", run.RunCodeHandler)
	r.POST("/api/submit/:id", run.SubmitSolutionHandler)

	r.POST("/signup", RegisterHandler)
	r.POST("/signin", SignInHandler)
	r.GET("/api/tasks", GetTasksHandler)
	r.GET("/api/tasks/:id", GetTaskByIDHandler)

	protected := r.Group("/api")
	protected.Use(auth.AuthRequired())
	protected.GET("/profile", ProfileHandler)
}
