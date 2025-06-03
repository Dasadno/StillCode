package rest

import (
	_ "net/http"

	"github.com/gin-gonic/gin"
)

/*
func GetHandler() {
	if g, err := db.Query("SELECT id, login, password") {

	}
}
*/

func LandingPage(c *gin.Context) {
	c.File("../../client/src/pages/landingPage.html")
}

func SignInPage(c *gin.Context) {
	c.File("../../client/src/pages/SignIn.html")
}

func SignUpPage(c *gin.Context) {
	c.File("../../client/src/pages/SignUp.html")
}
