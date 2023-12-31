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
func (u *UserRepository) GetById(id int) (*domain.User, error) {
	var user *domain.User
	err := u.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err 
	}
	return user, nil
}

// 	GetAll() ([]*User, error)
func (u *UserRepository) GetAll() ([]*domain.User, error) {
	var users []*domain.User
	err := u.db.Find(&users).Error
	return users, err
}

// Update(id int) (*User, error)
func (u *UserRepository) Update(newUser *domain.User) (*domain.User, error) {
	var user *domain.User
	err := u.db.Model(&user).Where("id = ?", newUser.ID).Updates(newUser).Find(&user).Error
	return user, err
}

// Delete(id int) (error)
func (u *UserRepository) Delete(id int) (*domain.User, error) {
	var user *domain.User
	err := u.db.Delete(&user, "id = ?", id).Find(&user).Error
	return user, err
}