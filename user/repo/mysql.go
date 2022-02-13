package mysql

import (
	"fmt"
	"time"
	"users/domain"

	"github.com/jinzhu/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

// NewMysqlUserRepository : create instance of repository
func NewMysqlUserRepository(con *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{con}
}
func (m *mysqlUserRepository) FetchUsers() (users []domain.User, err error) {
	m.db.Find(&users)
	return
}
func (m *mysqlUserRepository) GetUser(id int) (user domain.User, err error) {
	err = m.db.First(&user, id).Error
	return
}

func (m *mysqlUserRepository) CreateUser(user domain.User) (err error) {
	user.CreatedAt = time.Now()
	err = m.db.Create(&user).Error
	return
}

func (m *mysqlUserRepository) UpdateUser(user domain.User) (err error) {
	fmt.Printf("ashas")
	result := m.db.Model(&domain.User{}).Update(user).RowsAffected
	if result == 0 {
		err = domain.ErrUserNotFound
	}
	return
}

func (m *mysqlUserRepository) DeleteUser(id int) (err error) {
	result := m.db.Delete(&domain.User{}, id).RowsAffected
	if result == 0 {
		err = domain.ErrUserNotFound
	}
	return
}
