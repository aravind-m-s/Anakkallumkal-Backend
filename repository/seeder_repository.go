package repository

import (
	"github.com/aravind-m-s/anakallumkal-backend/domain"
	"gorm.io/gorm"
)

type SeederRepository interface {
	ShopSeeder() string
	ShopGet() ([]domain.Shop, error)
}

type seederDbStruct struct {
	DB *gorm.DB
}

func InitSeederRepo(db *gorm.DB) SeederRepository {
	return &seederDbStruct{DB: db}
}

func (s *seederDbStruct) ShopSeeder() string {

	shop1 := domain.Shop{
		Name:  "Anakallumkal Office Furnitures",
		Place: "Pala, Main Road",
		Image: "",
	}

	dbErr := s.DB.Create(&shop1).Error

	if dbErr != nil {
		return "Unable to Create shop 1 error:" + dbErr.Error()
	}

	shop2 := domain.Shop{
		Name:  "Anakallumkal Furnitures",
		Place: "Pala",
		Image: "",
	}

	dbErr = s.DB.Create(&shop2).Error

	if dbErr != nil {
		return "Unable to Create shop 2 error:" + dbErr.Error()
	}

	return "Success"

}

func (s *seederDbStruct) ShopGet() ([]domain.Shop, error) {

	var shops []domain.Shop

	dbErr := s.DB.Find(&shops).Error

	return shops, dbErr

}
