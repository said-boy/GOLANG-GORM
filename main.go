package main

import (
	"github.com/said-boy/golang-gorm/config/database"
	"github.com/said-boy/golang-gorm/domain"
	"github.com/said-boy/golang-gorm/repository"
	"github.com/said-boy/golang-gorm/usecase"
)

func main(){
	db := database.OpenConnection()

	a, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer a.Close()

	user := domain.User{
		Name: "Said",
		Gender: "male",
	}
	u := usecase.NewUserUsecase(repository.NewUserRepository(db))
	u.CreateUser(&user)

	user2 := domain.User{
		Name: "Cindy",
		Gender: "female",
	}
	u2 := usecase.NewUserUsecase(repository.NewUserRepository(db))
	u2.CreateUser(&user2)
}