package usecase

import (
	"github.com/said-boy/golang-gorm/domain"
	"github.com/said-boy/golang-gorm/repository"
)

type UserUsecase struct {
	UserRepository *repository.UserRepository
}

func NewUserUsecase(userRepository *repository.UserRepository ) *UserUsecase {
	return &UserUsecase{UserRepository: userRepository}
}

func (u *UserUsecase) CreateUser(user *domain.User) error {
	return u.UserRepository.Save(user)
}