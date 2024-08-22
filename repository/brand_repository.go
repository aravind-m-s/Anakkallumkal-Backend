package repository

import (
	"fmt"

	"github.com/aravind-m-s/anakallumkal-backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BrandRepoInterface interface {
	CreateBrand(name string, image string, shop uuid.UUID) (furniture domain.Brand, errorMsg error)
	DeleteBrand(id uuid.UUID) (errorMsg error)
	ListBrand() (furniture []domain.Brand, shops []domain.Shop, errorMsg error)
	UpdateBrand(id uuid.UUID, name string, image string, shop uuid.UUID) (furniture domain.Brand, errorMsg error)
}

type brandDbStruct struct {
	DB *gorm.DB
}

func InitBrandRepo(db *gorm.DB) BrandRepoInterface {
	return &brandDbStruct{DB: db}
}

func (f *brandDbStruct) CreateBrand(name string, image string, shop uuid.UUID) (brand domain.Brand, errorMsg error) {
	defer func() {
		if r := recover(); r != nil {
			errorMsg = r.(error)
		}
	}()

	var dbShop domain.Shop

	shopErr := f.DB.Find(&dbShop).Error

	if shopErr != nil {
		return brand, shopErr
	}

	if dbShop.Name == "" {
		return brand, fmt.Errorf("Invalid shop id")
	}

	brand.Name = name
	brand.Image = image
	brand.ShopID = shop

	dbErr := f.DB.Create(&brand).Error

	if dbErr != nil {
		return domain.Brand{}, dbErr
	}

	brand.Shop = dbShop

	return brand, nil
}

func (f *brandDbStruct) DeleteBrand(id uuid.UUID) (errorMsg error) {
	defer func() {
		if r := recover(); r != nil {
			errorMsg = r.(error)
		}
	}()

	tx := f.DB.Begin()

	var furniture []domain.Furniture

	dbErr := tx.Model(&domain.Furniture{}).Where("brand_id = ?", id).Find(&furniture).Error

	if dbErr != nil {
		tx.Rollback()
		return dbErr
	}

	for _, furn := range furniture {

		dbErr = tx.Delete(&furn).Error

		if dbErr != nil {
			tx.Rollback()
			return dbErr
		}

	}

	var brand domain.Brand

	dbErr = tx.Find(&brand).Error

	if dbErr != nil {
		tx.Rollback()
		return dbErr
	}

	dbErr = tx.Delete(&brand).Error

	if dbErr != nil {
		tx.Rollback()
		return dbErr
	}

	tx.Commit()

	return nil
}

func (f *brandDbStruct) ListBrand() (brands []domain.Brand, shops []domain.Shop, errorMsg error) {
	defer func() {
		if r := recover(); r != nil {
			errorMsg = r.(error)
		}
	}()

	dbErr := f.DB.Preload("Shop").Order("created_at DESC").Find(&brands).Error

	for index := range brands {
		dbErr = f.DB.Model(&domain.Furniture{}).Where("brand_id = ?", brands[index].ID).Count(&brands[index].Count).Error
	}

	dbErr = f.DB.Find(&shops).Error

	if dbErr != nil {
		return brands, shops, dbErr
	}

	if brands == nil {
		brands = []domain.Brand{}

	}

	return brands, shops, nil
}

func (f *brandDbStruct) UpdateBrand(id uuid.UUID, name string, image string, shop uuid.UUID) (brand domain.Brand, errorMsg error) {
	defer func() {
		if r := recover(); r != nil {
			errorMsg = r.(error)
		}
	}()

	dbErr := f.DB.Model(&brand).Where("id = ?", id).Find(&brand).Error

	if dbErr != nil || brand.Name == "" {
		return brand, fmt.Errorf("Invalid Brand Id")
	}

	brand.Name = name
	brand.ShopID = shop

	if image != "" {
		brand.Image = image
	}

	dbErr = f.DB.Updates(&brand).Error

	if dbErr != nil {
		return domain.Brand{}, dbErr
	}

	return brand, nil
}
