package http

import (
	"net/http"
	"strconv"
	"users/domain"

	"github.com/labstack/echo"
)

// UserHandler : http handler for users
type UserHandler struct {
	UserUsecase domain.UserUsecase
}

// NewUserHandler will initialize the users/ resources endpoint
func NewUserHandler(e *echo.Echo, us domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: us,
	}
	e.GET("/users", handler.FetchUsers)
	e.POST("/users", handler.CreateUser)
	e.GET("/users/:id", handler.GetUser)
	e.PUT("/users", handler.UpdateUser)
	e.DELETE("/users/:id", handler.DeleteUser)
}

// CreateUser : Creates a new user in the database
func (uh *UserHandler) CreateUser(c echo.Context) error {
	var user domain.User
	err := c.Echo().Binder.Bind(&user, c)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = uh.UserUsecase.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Success")
}

// FetchUsers : Gets all the user details from the database
func (uh *UserHandler) FetchUsers(c echo.Context) error {

	users, err := uh.UserUsecase.FetchUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

// GetUser : Gets a user details from the database
func (uh *UserHandler) GetUser(c echo.Context) error {

	idString := c.Param("id")
	var id int
	if len(idString) == 0 {
		return c.JSON(http.StatusOK, domain.MIssingUserID)
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusOK, domain.InvalidUserID)
	}
	userResponse, err := uh.UserUsecase.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, userResponse)
}

// UpdateUser : Updates a user in the database
func (uh *UserHandler) UpdateUser(c echo.Context) error {
	var user domain.User
	err := c.Echo().Binder.Bind(&user, c)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = uh.UserUsecase.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "User Details updated successfully")
}

// DeleteUser : deletes a user from the database
func (uh *UserHandler) DeleteUser(c echo.Context) error {
	var user domain.User
	err := c.Echo().Binder.Bind(&user, c)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	idString := c.Param("id")
	var id int
	if len(idString) == 0 {
		return c.JSON(http.StatusOK, domain.MIssingUserID)
	}

	id, err = strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusOK, domain.InvalidUserID)
	}

	err = uh.UserUsecase.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Successfully DeleteUserd User ID:"+idString)
}
