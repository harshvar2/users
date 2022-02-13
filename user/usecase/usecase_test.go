package usecase

import (
	"testing"
	"users/domain"
	"users/domain/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchUsers(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	t.Run("Test1", func(t *testing.T) {
		mockUserRepo.On("FetchUsers").Return([]domain.User{}, nil).Once()
		u := NewUserUsecase(mockUserRepo)
		result, err := u.FetchUsers()
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.IsType(t, []domain.User{}, result)
	})
}

func TestGetUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	t.Run("Test1", func(t *testing.T) {
		userID := 1
		mockUserRepo.On("GetUser", mock.AnythingOfType("int")).Return(domain.User{}, nil).Once()
		u := NewUserUsecase(mockUserRepo)
		result, err := u.GetUser(userID)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.IsType(t, domain.User{}, result)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestCreateUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	t.Run("Test1", func(t *testing.T) {
		user := domain.User{Name: "User", Address: "localhost", DateOfBirth: "29-09-2014"}
		mockUserRepo.On("CreateUser", mock.AnythingOfType("domain.User")).Return(nil).Once()
		u := NewUserUsecase(mockUserRepo)
		err := u.CreateUser(user)
		assert.Nil(t, err)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	t.Run("Test1", func(t *testing.T) {
		user := domain.User{Name: "User", Address: "localhost", DateOfBirth: "29-09-2014"}
		mockUserRepo.On("UpdateUser", mock.AnythingOfType("domain.User")).Return(nil).Once()
		u := NewUserUsecase(mockUserRepo)
		err := u.UpdateUser(user)
		assert.Nil(t, err)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	t.Run("Test1", func(t *testing.T) {
		userID := 1
		mockUserRepo.On("DeleteUser", mock.AnythingOfType("int")).Return(nil).Once()
		u := NewUserUsecase(mockUserRepo)
		err := u.DeleteUser(userID)
		assert.Nil(t, err)
		mockUserRepo.AssertExpectations(t)
	})
}
