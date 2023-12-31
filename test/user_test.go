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

func TestGetAllUserSuccessful(t *testing.T) {
	db := database.OpenConnection()

	a, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer a.Close()

	userUsecase := usecase.NewUserUsecase(repository.NewUserRepository(db))

	users, err := userUsecase.GetAllUser()
	
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))
}

func TestUpdateUserSuccessful(t *testing.T) {
	db := database.OpenConnection()

	a, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer a.Close()

	userUsecase := usecase.NewUserUsecase(repository.NewUserRepository(db))

	updateUser := domain.User{
		ID: "2",
		Name: "Joko",
	}
	user, err := userUsecase.UpdateUser(&updateUser)
	
	assert.Nil(t, err)
	assert.Equal(t, "Joko", user.Name)
}

func TestDeleteUserSuccessful(t *testing.T) {
	db := database.OpenConnection()

	a, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer a.Close()

	userUsecase := usecase.NewUserUsecase(repository.NewUserRepository(db))

	user, err := userUsecase.DeleteUser(34)
	
	assert.Nil(t, err)
	assert.Equal(t, 0, len(user.ID))
}
