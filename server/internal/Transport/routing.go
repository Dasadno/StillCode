package Transport

import (
	_ "StillCode/server/internal/Transport/rest"

	"github.com/gin-gonic/gin"
)

func SetupWebRoutes(r *gin.Engine) {

	r.GET("/", LandingPage)
	r.GET("/signin", SignInPage)
	r.GET("/signup", SignUpPage)
	r.GET("/profile", ProfilePage)
}

func ProfilePage(c *gin.Context) {
	c.File("../../client/src/pages/profile.html")
}

func LandingPage(c *gin.Context) {
	c.File("../../client/src/pages/landingPage.html")
}

func SignInPage(c *gin.Context) {
	c.File("../../client/src/pages/SignIn.html")
}

func SignUpPage(c *gin.Context) {
	c.File("../../client/src/pages/SignUp.html")
}
