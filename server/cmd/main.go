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

	r.Static("/js", "../../client/src/script")
	r.Static("/css", "../../client/src/styles")
	r.LoadHTMLGlob("../../server/templates/*.html")

	r.POST("/register", rest.RegisterHandler)
	r.GET("/signin", rest.SignUpHandler)

	Transport.SetupWebRoutes(r)

	r.Run(":8080")

}
