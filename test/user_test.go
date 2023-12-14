package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/said-boy/golang-gorm/config/database"
	"github.com/said-boy/golang-gorm/domain"
	"github.com/said-boy/golang-gorm/repository"
	"github.com/said-boy/golang-gorm/usecase"
)

func TestCreateUserSuccessful(t *testing.T) {
	db := database.OpenConnection()

	a, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer a.Close()

	userUsecase := usecase.NewUserUsecase(repository.NewUserRepository(db))

	user := domain.User{
		Name:   "Said",
		Gender: "male",
	}
	err = userUsecase.CreateUser(&user)

	assert.Nil(t, err)
}

func TestCreateUserFail(t *testing.T) {
	db := database.OpenConnection()

	a, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer a.Close()

	userUsecase := usecase.NewUserUsecase(repository.NewUserRepository(db))

	user := domain.User{
		Gender: "male",
	}
	err = userUsecase.CreateUser(&user)

	assert.NotNil(t, err)
	assert.Equal(t, "Field Name in Entity User is required", err.Error())
}

func TestGetUserByIdSuccessful(t *testing.T) {
	db := database.OpenConnection()

	a, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer a.Close()

	userUsecase := usecase.NewUserUsecase(repository.NewUserRepository(db))

	user, err := userUsecase.GetUserById(1)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(user.ID))
}

func TestGetUserByIdFail(t *testing.T) {
	db := database.OpenConnection()

	a, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer a.Close()

	userUsecase := usecase.NewUserUsecase(repository.NewUserRepository(db))

	user, err := userUsecase.GetUserById(4)

	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "record not found", err.Error())
}
