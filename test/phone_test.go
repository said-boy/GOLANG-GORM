package test

import (
	"testing"

	"github.com/said-boy/golang-gorm/config/database"
	"github.com/said-boy/golang-gorm/domain"
	"github.com/said-boy/golang-gorm/repository"
	"github.com/said-boy/golang-gorm/usecase"
	"github.com/stretchr/testify/assert"
)

func TestSavePhone(t *testing.T) {
	db := database.OpenConnection()

	a, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer a.Close()

	phoneUsecase := usecase.NewPhoneUsecase(repository.NewPhoneRepository(db))

	phoneUser1 := domain.Phone{
		Handphone: "+6289582955000",
		UserId:    1, // id user said
	}
	err = phoneUsecase.SavePhone(&phoneUser1)

	assert.Nil(t, err)

}

