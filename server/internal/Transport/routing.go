package Transport

import (
	"StillCode/server/internal/Transport/rest"

	"github.com/gin-gonic/gin"
)

func SetupWebRoutes(r *gin.Engine) {

	r.GET("/", rest.LandingPage)

	r.GET("/signin", rest.SignInPage)

	r.GET("/signup", rest.SignUpPage)

}
