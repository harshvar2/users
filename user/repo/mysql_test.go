package mysql_test

import (
	"testing"
	"time"
	"users/domain"
	mysql "users/user/repo"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var userColumn = []string{"id", "name", "date_of_birth", "address", "created_at", "updated_at"}

func TestFetchUsers(t *testing.T) {
	DB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	db, err := gorm.Open("mysql", DB)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	currentTime := time.Now()

	rows := sqlmock.NewRows(userColumn).
		AddRow(1, "User1", "19-09-1999", "INDIA", currentTime, currentTime).
		AddRow(2, "User2", "29-09-1999", "INDIA", currentTime, currentTime)

	mock.ExpectQuery("^SELECT (.+) FROM `users`").WillReturnRows(rows)

	mysqlRepo := mysql.NewMysqlUserRepository(db)

	users, err := mysqlRepo.FetchUsers()

	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, len(users), 2)
	assert.Nil(t, err)
}

func TestGetUser(t *testing.T) {
	DB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	db, err := gorm.Open("mysql", DB)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	currentTime := time.Now()

	row := sqlmock.NewRows(userColumn).
		AddRow(1, "User1", "19-09-1999", "INDIA", currentTime, currentTime)

	mock.ExpectQuery("^SELECT (.+) FROM `users` WHERE .*").WillReturnRows(row)

	userID := 1
	mysqlRepo := mysql.NewMysqlUserRepository(db)
	user, err := mysqlRepo.GetUser(userID)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, 1)
	assert.NotNil(t, user)
}

func TestCreateUser(t *testing.T) {
	DB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	db, err := gorm.Open("mysql", DB)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `users` .*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	currentTime := time.Now()

	newUser := domain.User{ID: 1, Name: "User", Address: "INDIA", CreatedAt: currentTime}

	mysqlRepo := mysql.NewMysqlUserRepository(db)
	err = mysqlRepo.CreateUser(newUser)
	assert.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	DB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	db, err := gorm.Open("mysql", DB)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE `users` SET .*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	currentTime := time.Now()

	newUser := domain.User{ID: 1, Name: "User", Address: "INDIA", CreatedAt: currentTime}

	mysqlRepo := mysql.NewMysqlUserRepository(db)
	err = mysqlRepo.UpdateUser(newUser)
	assert.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	DB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	db, err := gorm.Open("mysql", DB)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("^DELETE FROM `users` WHERE .*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	userID := 1

	mysqlRepo := mysql.NewMysqlUserRepository(db)
	err = mysqlRepo.DeleteUser(userID)
	assert.NoError(t, err)
}
