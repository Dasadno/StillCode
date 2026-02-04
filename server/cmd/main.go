package main

import (
	"StillCode/server/internal/db"
	"StillCode/server/internal/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// Initialize database
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Run migrations
	if err := db.RunMigrations(); err != nil {
		log.Printf("Warning: migrations failed: %v", err)
	}

	// Setup router with API routes
	r := router.SetupRouter()

	// Setup static files and web routes
	router.SetupStaticFiles(r)
	router.LoadTemplates(r)
	router.SetupWebRoutes(r)

	// Get port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Check for TLS certificates
	certPath := os.Getenv("CERT_PATH")
	keyPath := os.Getenv("KEY_PATH")

	// If no certs provided, check default locations
	if certPath == "" {
		certPath = "./certificate.crt"
	}
	if keyPath == "" {
		keyPath = "./private.key"
	}

	// Check if certificate files exist
	_, certErr := os.Stat(certPath)
	_, keyErr := os.Stat(keyPath)

	if certErr == nil && keyErr == nil {
		// Run with TLS
		log.Printf("Starting HTTPS server on :%s", port)
		if err := r.RunTLS(":"+port, certPath, keyPath); err != nil {
			log.Fatalf("Failed to start HTTPS server: %v", err)
		}
	} else {
		// Run without TLS (development mode)
		log.Printf("Starting HTTP server on :%s (no TLS certificates found)", port)
		if err := r.Run(":" + port); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}
}
