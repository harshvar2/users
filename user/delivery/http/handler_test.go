package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"users/domain"
	"users/domain/mocks"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestFetchUsers(t *testing.T) {
	var mockResponse domain.User
	err := faker.FakeData(&mockResponse)
	assert.NoError(t, err)

	e := echo.New()

	mockUCase := new(mocks.UserUsecase)
	mockUCase.On("FetchUsers").Return([]domain.User{}, nil)

	req, err := http.NewRequest(echo.GET, "/users", strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	handler := UserHandler{
		UserUsecase: mockUCase,
	}
	err = handler.FetchUsers(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotNil(t, rec.Body)
	mockUCase.AssertExpectations(t)
}

func TestGetUser(t *testing.T) {
	var mockResponse domain.User
	err := faker.FakeData(&mockResponse)
	assert.NoError(t, err)

	e := echo.New()

	mockUCase := new(mocks.UserUsecase)
	mockUCase.On("GetUser", mock.AnythingOfType("int")).Return(domain.User{}, nil)

	req, err := http.NewRequest(echo.GET, "/users/:id", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	handler := UserHandler{
		UserUsecase: mockUCase,
	}
	err = handler.GetUser(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotNil(t, rec.Body)
	mockUCase.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	var mockResponse domain.User
	err := faker.FakeData(&mockResponse)
	assert.NoError(t, err)

	e := echo.New()

	mockUCase := new(mocks.UserUsecase)
	mockUCase.On("CreateUser", mock.AnythingOfType("domain.User")).Return(nil)

	var jsonStr = []byte(`{
        "id": 121,
        "name": "User",
        "dob": "19-09-1999",
        "address": "india",
        "createdAt": "2022-02-13T23:41:12+05:30",
        "updatedAt": "2022-02-13T23:41:12+05:30"
    }`)
	reqBody := bytes.NewBuffer(jsonStr)

	req, err := http.NewRequest(echo.POST, "/users", reqBody)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	handler := UserHandler{
		UserUsecase: mockUCase,
	}
	err = handler.CreateUser(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotNil(t, rec.Body)
	mockUCase.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	var mockResponse domain.User
	err := faker.FakeData(&mockResponse)
	assert.NoError(t, err)

	e := echo.New()

	mockUCase := new(mocks.UserUsecase)
	mockUCase.On("UpdateUser", mock.AnythingOfType("domain.User")).Return(nil)

	var jsonStr = []byte(`{
        "id": 121,
        "name": "User",
        "dob": "19-09-1999",
        "address": "india"
    }`)
	reqBody := bytes.NewBuffer(jsonStr)

	req, err := http.NewRequest(echo.PUT, "/users", reqBody)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")
	handler := UserHandler{
		UserUsecase: mockUCase,
	}
	err = handler.UpdateUser(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotNil(t, rec.Body)
	mockUCase.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	var mockResponse domain.User
	err := faker.FakeData(&mockResponse)
	assert.NoError(t, err)

	e := echo.New()

	mockUCase := new(mocks.UserUsecase)
	mockUCase.On("DeleteUser", mock.AnythingOfType("int")).Return(nil)

	var jsonStr = []byte(`{
        "id": 121,
        "name": "User",
        "dob": "19-09-1999",
        "address": "india"
    }`)
	reqBody := bytes.NewBuffer(jsonStr)

	req, err := http.NewRequest(echo.DELETE, "/users/:id", reqBody)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	handler := UserHandler{
		UserUsecase: mockUCase,
	}
	err = handler.DeleteUser(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotNil(t, rec.Body)
	mockUCase.AssertExpectations(t)
}
