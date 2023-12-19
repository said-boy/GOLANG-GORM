package domain

type User struct {
	ID     string `gorm:"primaryKey;autoIncrement;column:id;"`
	Name   string `gorm:"column:name;" validate:"required"`
	Gender string `gorm:"column:gender;"`
	Phone []Phone `gorm:"foreignKey:user_id;references:id"`
}

func (u *User) TableName() string {
	return "users"
}
