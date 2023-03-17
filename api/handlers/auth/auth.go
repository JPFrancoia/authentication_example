package auth

import (
	"fmt"
	"net/http"
	"os"

	data_registry "local/auth_example/api/data_registy"
	ent "local/auth_example/api/entities"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
)

func init() {
	facebookProvider := facebook.New(
		os.Getenv("FACEBOOK_CLIENT_ID"),
		os.Getenv("FACEBOOK_CLIENT_SECRET"),
		os.Getenv("AUTH_REDIRECT_URL"),
	)
	goth.UseProviders(facebookProvider)
}

func Login(c *gin.Context) {
	// Insert provider into context
	// https://github.com/markbates/goth/issues/411#issuecomment-891037676
	q := c.Request.URL.Query()

	q.Add("provider", c.Param("provider"))
	c.Request.URL.RawQuery = q.Encode()

	fmt.Println("Starting auth flow")
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func AuthCallback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)

	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	ourUser := ent.GothUser(user).ToUser()

	// Write the new user to database.
	// If the user already exists, we'll refresh some minor info e.g
	// first/last name, but nothing else will change, they'll keep their
	// existing user_id
	err = data_registry.UpsertUser(ourUser)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusFound, "/")
}
