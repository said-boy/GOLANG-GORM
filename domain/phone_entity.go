package domain

import (
	"github.com/said-boy/golang-gorm/validate"
)

type Phone struct {
	Handphone string `gorm:"column:handphone;primaryKey;" validate:"required,e164"`
	UserId    int    `gorm:"column:user_id;primaryKey;"`
}

func (p *Phone) TableName() string {
	return "phone"
}

func (p *Phone) Validate(phone *Phone) error {
	return validate.Validate.Struct(phone)
}
