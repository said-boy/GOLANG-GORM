package domain

type UserRepository interface {
	Save(user *User) error
	GetById(id int) (*User, error)
	GetAll() ([]*User, error)
	Update(id int) (*User, error)
	// Delete(id int) (error)
}