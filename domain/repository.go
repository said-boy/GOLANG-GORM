package domain

type UserRepository interface {
	Save(user *User) error
	GetById(id int) (*User, error)
	GetAll() ([]*User, error)
	Update(newUser *User) (*User, error)
	Delete(id int) (*User, error)
}