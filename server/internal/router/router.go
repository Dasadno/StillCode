package router

import (
	"StillCode/server/internal/auth"
	"StillCode/server/internal/handlers"
	"StillCode/server/internal/middleware"
	"StillCode/server/internal/services"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures and returns the Gin router with all routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Apply global middleware
	r.Use(middleware.CORS())

	// Initialize services
	authService := services.NewAuthService()
	taskService := services.NewTaskService()
	userService := services.NewUserService()
	submissionService := services.NewSubmissionService(taskService)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	taskHandler := handlers.NewTaskHandler(taskService)
	userHandler := handlers.NewUserHandler(userService)
	submissionHandler := handlers.NewSubmissionHandler(submissionService)

	// Rate limiter: 10 requests per minute for code execution
	runLimiter := middleware.NewRateLimiter(10, time.Minute)

	// API routes group
	api := r.Group("/api")
	{
		// Auth routes - public
		authRoutes := api.Group("/auth")
		{
			authRoutes.POST("/signup", authHandler.SignUp)
			authRoutes.POST("/signin", authHandler.SignIn)
		}

		// Task routes - public
		api.GET("/tasks", taskHandler.GetTasks)
		api.GET("/tasks/:id", taskHandler.GetTaskByID)

		// Protected routes - require authentication
		protected := api.Group("")
		protected.Use(auth.AuthRequired())
		{
			// Profile route
			protected.GET("/profile", userHandler.GetProfile)

			// Code execution routes - protected and rate limited
			protected.POST("/run", runLimiter.Middleware(), submissionHandler.RunCode)
			protected.POST("/submit/:id", runLimiter.Middleware(), submissionHandler.SubmitSolution)
		}
	}

	return r
}

// SetupWebRoutes configures routes for serving static HTML pages
func SetupWebRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.File("./client/src/pages/landing.html")
	})
	r.GET("/signin", func(c *gin.Context) {
		c.File("./client/src/pages/signin.html")
	})
	r.GET("/signup", func(c *gin.Context) {
		c.File("./client/src/pages/signup.html")
	})
	r.GET("/profile", func(c *gin.Context) {
		c.File("./client/src/pages/profile.html")
	})
	r.GET("/problems", func(c *gin.Context) {
		c.File("./client/src/pages/tasks.html")
	})
	r.GET("/task/:id", func(c *gin.Context) {
		c.File("./client/src/pages/task.html")
	})

	// 404 handler
	r.NoRoute(func(c *gin.Context) {
		c.File("./client/src/pages/404.html")
	})
}

// SetupStaticFiles configures static file serving
func SetupStaticFiles(r *gin.Engine) {
	r.Static("/src", "./client/src")
}

// LoadTemplates loads HTML templates (kept for backward compatibility)
func LoadTemplates(r *gin.Engine) {
	// Templates no longer used - serving static files directly
}
