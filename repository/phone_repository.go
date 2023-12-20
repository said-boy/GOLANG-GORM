package repository

import (
	"github.com/said-boy/golang-gorm/domain"
	"gorm.io/gorm"
)

type PhoneRepository struct {
	db *gorm.DB
}

func NewPhoneRepository(db *gorm.DB) *PhoneRepository {
	return &PhoneRepository{db: db}
}

// Save(*Phone) error
func (p *PhoneRepository) Save(phone *domain.Phone) error {
	return p.db.Save(phone).Error
}

// Update(*Phone) error
func (p *PhoneRepository) Update(phoneOld *domain.Phone, phoneNew *domain.Phone) error {
	return p.db.Model(domain.Phone{}).Where("handphone = ?", phoneOld.Handphone).Where("user_id = ?", phoneOld.UserId).Updates(phoneNew).Error
}

// Delete(id int) error
func (p *PhoneRepository) Delete(id int) error {
	var phone *domain.Phone
	return p.db.Delete(&phone, "id = ?", id).Error
}
