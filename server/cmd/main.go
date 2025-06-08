package main

import (
	"StillCode/server/internal/Transport"
	"StillCode/server/internal/Transport/rest"
	_ "StillCode/server/internal/Transport/rest"
	"StillCode/server/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	db.InitDb()

	r.POST("/register", rest.RegisterHandler)

	r.Static("/js", "../../client/src/script")
	r.Static("/css", "../../client/src/styles")

	r.LoadHTMLGlob("../../server/templates/*.html")

	Transport.SetupWebRoutes(r)

	r.Run(":8080")

}
