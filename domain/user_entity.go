package domain

type User struct {
	ID     string `gorm:"primaryKey;autoIncrement;column:id;"`
	Name   string `gorm:"column:name;"`
	Gender string `gorm:"column:gender;"`
}

func (u *User) TableName() string {
	return "users"
}
