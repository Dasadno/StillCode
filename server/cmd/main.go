package main

import (
	"StillCode/server/internal/Transport"
	"StillCode/server/internal/Transport/rest"
	"StillCode/server/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	db.InitDb()
	defer db.DB.Close()
	r.Static("/js", "../../client/src/script")
	r.Static("/css", "../../client/src/styles")
	r.LoadHTMLGlob("../../server/templates/*.html")

	rest.SetupHandlers(r)
	Transport.SetupWebRoutes(r)
	CertificatePath := "C:\\Users\\1\\sc\\certificate.crt"
	KeyPath := "C:\\Users\\1\\sc\\private.key"

	err := r.RunTLS(":8080", CertificatePath, KeyPath)
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}

}
