package rest

import (
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func SetupHandlers(r *gin.Engine) {

	r.POST("/signup", RegisterHandler)

}
