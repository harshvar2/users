package domain

import (
	"time"
)

// User struct for 'users' table
type User struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	DateOfBirth string    `json:"dob"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// UserRepository represents a usecase interface for User
type UserRepository interface {
	FetchUsers() ([]User, error)
	CreateUser(user User) (err error)
	UpdateUser(user User) (err error)
	DeleteUser(id int) (err error)
	GetUser(id int) (res User, err error)
}

// UserUsecase represents a usecase interface for User
type UserUsecase interface {
	FetchUsers() ([]User, error)
	CreateUser(user User) (err error)
	UpdateUser(user User) (err error)
	DeleteUser(id int) (err error)
	GetUser(id int) (res User, err error)
}
