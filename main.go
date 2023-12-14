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

	u := usecase.NewUserUsecase(repository.NewUserRepository(db))

	// user := domain.User{
	// 	Name: "Said",
	// 	Gender: "male",
	// }
	// u.CreateUser(&user)

	// user2 := domain.User{
	// 	Name: "Cindy",
	// 	Gender: "female",
	// }
	// u.CreateUser(&user2)

	user3 := domain.User{
		Gender: "male",
	}

	err = u.CreateUser(&user3)
	if err != nil {
		panic(err)
	}
}