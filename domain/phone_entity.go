package domain

import (
	"github.com/said-boy/golang-gorm/validate"
)

type Phone struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement"`
	Handphone string `gorm:"column:handphone;" validate:"required,e164"`
	UserId    int   `gorm:"column:user_id"`
}

func (p *Phone) TableName() string {
	return "phone"
}

func (p *Phone) Validate(phone *Phone) error {
	return validate.Validate.Struct(phone)
}