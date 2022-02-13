package usecase

import (
	"users/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

// NewUserUsecase represents a new instance for User Interface
func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
func (u *userUsecase) FetchUsers() (users []domain.User, err error) {
	users, err = u.userRepo.FetchUsers()
	return
}
func (u *userUsecase) GetUser(id int) (userResponse domain.User, err error) {
	userResponse, err = u.userRepo.GetUser(id)
	return
}
func (u *userUsecase) CreateUser(user domain.User) (err error) {
	err = u.userRepo.CreateUser(user)
	return
}

func (u *userUsecase) UpdateUser(user domain.User) (err error) {
	err = u.userRepo.UpdateUser(user)
	return
}
func (u *userUsecase) DeleteUser(id int) (err error) {
	err = u.userRepo.DeleteUser(id)
	return
}
