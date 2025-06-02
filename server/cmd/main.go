package main

import (
	_ "StillCode/server/internal/Transport/rest"
	"StillCode/server/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.StaticFile("/", "../../client/src/assets/index.html")

	db.InitDb()

}
