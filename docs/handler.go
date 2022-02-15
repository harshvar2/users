package http

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type documentationCredentials struct {
	Credentials []documentationCredential
}

type documentationCredential struct {
	username string
	password string
}

// The documentation route is password protected using basic auth. The credentials are stored in an environment variable.
func getDocumentationCredentials() (response documentationCredentials, err error) {

	_credentials := viper.GetString("DOCUMENTATION_CREDENTIALS")
	if len(_credentials) == 0 {
		err = errors.New("missingDocumentationCredentials")
		return
	}

	// If the credentials are not of the form "user1:password1;user2:password2" then throw an error
	credentials := strings.Split(_credentials, ";")

	for _, credential := range credentials {
		c := strings.Split(credential, ":")
		if len(c) != 2 {
			err = errors.New("invalidDocumentationCredtialsFormat")
			return
		}
		response.Credentials = append(response.Credentials, documentationCredential{
			username: c[0],
			password: c[1],
		})
	}
	return
}

// isValid checks if the given credntials are valid or not.
func (c *documentationCredential) isValid(existing documentationCredentials) (response bool) {
	for _, e := range existing.Credentials {
		if e.username == c.username && e.password == c.password {
			return true
		}
	}
	return
}

// initialiseAuth is a wrapper for the authorisation which will be used to limit access of documentation
func initialiseAuth(d *echo.Group) (err error) {

	existingCredentials, err := getDocumentationCredentials()
	if err != nil {
		return
	}

	d.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		credential := documentationCredential{
			username: username,
			password: password,
		}
		if credential.isValid(existingCredentials) {
			return true, nil
		}
		return false, nil
	}))

	return
}

// NewDocumentation will initialize the endpoints and their respective handlers and static pages for documentation
func NewDocumentation(e *echo.Echo, d *echo.Group) {

	err := initialiseAuth(d)
	if err != nil {
		e.Logger.Fatal(err)
	}
	d.Static("/docs", "docs")
	d.Static("/databaseSchema", "docs/database")
}
