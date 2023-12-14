package usecase

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/said-boy/golang-gorm/domain"
	"github.com/said-boy/golang-gorm/repository"
)

type UserUsecase struct {
	UserRepository *repository.UserRepository
	Validator *validator.Validate
}

func NewUserUsecase(userRepository *repository.UserRepository ) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
		Validator: validator.New(),
	}
}

func (u *UserUsecase) CreateUser(user *domain.User) error {
	err := u.Validator.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New("Field " + err.Field() + " in Entity User is " + err.Tag())
		}
	}
	u.UserRepository.Save(user)
	return nil
}