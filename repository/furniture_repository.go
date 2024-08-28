package repository

import (
	"fmt"

	"github.com/aravind-m-s/anakallumkal-backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FurnitureRepoInterface interface {
	CreateFurniture(name string, image string, productNo string, brand uuid.UUID, stock int, price int) (furniture domain.Furniture, errorMsg error)
	DeleteFurniture(id uuid.UUID) (errorMsg error)
	ListFurniture(id uuid.UUID, query string) (furniture []domain.Furniture, errorMsg error)
	UpdateFurniture(id uuid.UUID, name string, image string, productNo string, brand uuid.UUID, stock int, price int) (furniture domain.Furniture, errorMsg error)
	ExportFurniture(id uuid.UUID) (furnitures []domain.Furniture, errorMsg error)
}

type furnitureDbStruct struct {
	DB *gorm.DB
}

func InitFurnitureRepo(db *gorm.DB) FurnitureRepoInterface {
	return &furnitureDbStruct{DB: db}
}

func (f *furnitureDbStruct) CreateFurniture(name string, image string, productNo string, brand uuid.UUID, stock int, price int) (furniture domain.Furniture, errorMsg error) {
	defer func() {
		if r := recover(); r != nil {
			errorMsg = r.(error)
		}
	}()

	var dbBrand domain.Brand

	brandErr := f.DB.Find(&dbBrand).Error

	if brandErr != nil {
		return furniture, brandErr
	}

	if dbBrand.Name == "" {
		return furniture, fmt.Errorf("Invalid brand id")
	}

	furniture = domain.Furniture{
		Name:      name,
		Image:     image,
		ProductNo: productNo,
		BrandID:   brand,
		Stock:     stock,
		Price:     price,
	}

	dbErr := f.DB.Create(&furniture).Error

	if dbErr != nil {
		return domain.Furniture{}, dbErr
	}

	furniture.Brand = dbBrand

	return furniture, nil
}

func (f *furnitureDbStruct) DeleteFurniture(id uuid.UUID) (errorMsg error) {
	defer func() {
		if r := recover(); r != nil {
			errorMsg = r.(error)
			fmt.Printf("errorMsg.Error(): %v\n", errorMsg.Error())
		}
	}()

	tx := f.DB.Begin()

	fmt.Printf("id: %v\n", id)

	if tx == nil {
		fmt.Println("TX is nil")
		return
	}

	var furniture domain.Furniture

	dbErr := tx.Model(&domain.Furniture{}).Where("id = ?", id).First(&furniture).Error

	if dbErr != nil {
		fmt.Printf("dbErr.Error(): %v\n", dbErr.Error())
		tx.Rollback()
		return dbErr
	}

	dbErr = tx.Delete(&furniture).Error

	if dbErr != nil {
		tx.Rollback()
		fmt.Printf("dbErr.Error(): %v\n", dbErr.Error())
		return dbErr
	}

	tx.Commit()

	return nil
}

func (f *furnitureDbStruct) ExportFurniture(id uuid.UUID) (furnitures []domain.Furniture, errorMsg error) {
	defer func() {
		if r := recover(); r != nil {
			errorMsg = r.(error)
			fmt.Printf("errorMsg.Error(): %v\n", errorMsg.Error())
		}
	}()

	dbErr := f.DB.Preload("Brand").Where("brand_id = ?", id).Find(&furnitures).Error

	if dbErr != nil {
		return furnitures, dbErr
	}

	if len(furnitures) == 0 {
		furnitures = []domain.Furniture{}

	}

	return furnitures, nil
}

func (f *furnitureDbStruct) ListFurniture(id uuid.UUID, query string) (furniture []domain.Furniture, errorMsg error) {
	defer func() {
		if r := recover(); r != nil {
			errorMsg = r.(error)
		}
	}()

	var brand domain.Brand

	brandErr := f.DB.Model(&brand).Where("id = ?", id).Find(&brand).Error

	if brandErr != nil {
		return furniture, brandErr
	}

	if brand.Name == "" {
		return furniture, fmt.Errorf("No brand with the given id found")
	}

	if brand.Name == "" {
		return furniture, fmt.Errorf("Invalid brand id")
	}

	dbQuery := f.DB.Preload("Brand").Where("brand_id = ?", id)

	if query != "" {
		dbQuery = dbQuery.Where("name ILIKE ?", "%"+query+"%")
	}

	dbErr := dbQuery.Find(&furniture).Error

	if dbErr != nil {
		return furniture, dbErr
	}

	if furniture == nil {
		furniture = []domain.Furniture{}
	}

	return furniture, nil
}

func (f *furnitureDbStruct) UpdateFurniture(id uuid.UUID, name string, image string, productNo string, brand uuid.UUID, stock int, price int) (furniture domain.Furniture, errorMsg error) {
	defer func() {
		if r := recover(); r != nil {
			errorMsg = r.(error)
		}
	}()

	dbErr := f.DB.Model(&furniture).Where("id = ?", id).Find(&furniture).Error

	if dbErr != nil || furniture.Name == "" {
		return furniture, fmt.Errorf("Invalid Furniture id")
	}

	furniture.Name = name
	furniture.ProductNo = productNo
	furniture.BrandID = brand
	furniture.Stock = stock
	furniture.Price = price

	if image != "" {
		furniture.Image = image
	}

	dbErr = f.DB.Updates(&furniture).Error

	if dbErr != nil {
		return domain.Furniture{}, dbErr
	}

	return furniture, nil
}
