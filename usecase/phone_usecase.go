package usecase

import (
	"github.com/said-boy/golang-gorm/domain"
	"github.com/said-boy/golang-gorm/repository"
)

type PhoneUsecase struct {
	PhoneRepository *repository.PhoneRepository
	PhoneDomain *domain.Phone
}

func NewPhoneUsecase(phoneRepository *repository.PhoneRepository ) *PhoneUsecase {
	return &PhoneUsecase{
		PhoneRepository: phoneRepository,
	}
}

func (p *PhoneUsecase) SavePhone(phone *domain.Phone) error {
	err := p.PhoneDomain.Validate(phone)
	if err != nil {
		return err
	}
	return p.PhoneRepository.Save(phone)
}

func (p *PhoneUsecase) UpdatePhone(phone *domain.Phone) error {
	err := p.PhoneDomain.Validate(phone)
	if err != nil {
		return err
	}
	return p.PhoneRepository.Update(phone)
}

func (p *PhoneUsecase) DeletePhone(id int) error {
	return p.PhoneRepository.Delete(id)
}