package repository

import (
	"github.com/said-boy/golang-gorm/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// 	Save(user *User) error
func (u *UserRepository) Save(user *domain.User) error {
	return u.db.Save(user).Error
}

// 	GetById(id int) (*User, error)


// 	GetAll() ([]*User, error)