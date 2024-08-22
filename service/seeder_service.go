package service

import (
	"github.com/aravind-m-s/anakallumkal-backend/domain"
	"github.com/aravind-m-s/anakallumkal-backend/repository"
)

type SeederService interface {
	ShopSeeder() string
	ShopGet() ([]domain.Shop, error)
}

type seederServiceStruct struct {
	repo repository.SeederRepository
}

func InitSeederService(repo repository.SeederRepository) SeederService {
	return &seederServiceStruct{repo: repo}
}

func (s *seederServiceStruct) ShopSeeder() string {

	return s.repo.ShopSeeder()

}

func (s *seederServiceStruct) ShopGet() ([]domain.Shop, error) {
	return s.repo.ShopGet()
}
