package repository

import (
	"fmt"

	"github.com/aravind-m-s/anakallumkal-backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category *domain.Category) error
	CreateSubCategory(category *domain.SubCategory) error
	UpdateCategory(category *domain.Category) error
	UpdateSubCategory(category *domain.SubCategory) error
	DeleteCategory(id uuid.UUID) error
	DeleteSubCategory(id uuid.UUID) error
	ListCategory() ([]domain.CategoryResponse, error)
}

type categoryDbStruct struct {
	DB *gorm.DB
}

func InitCategoryRepo(db *gorm.DB) CategoryRepository {
	return &categoryDbStruct{DB: db}
}

func (c *categoryDbStruct) CreateCategory(category *domain.Category) (errorMsg error) {
	defer func() {
		if r := recover(); r != nil {
			errorMsg = r.(error)
		}
	}()

	dbErr := c.DB.Preload("Category").Create(&category).Error

	if dbErr != nil {
		return dbErr
	}

	return nil
}

func (c *categoryDbStruct) CreateSubCategory(category *domain.SubCategory) (errorMsg error) {
	defer func() {
		if r := recover(); r != nil {
			errorMsg = r.(error)
		}
	}()

	var dbCat domain.Category

	dbErr := c.DB.Model(&domain.Category{}).Where("id = ?", category.CategoryID).Find(&dbCat).Error

	if dbErr != nil {
		return dbErr
	}

	if dbCat.Name == "" {
		return fmt.Errorf("category does not exist")
	}

	dbErr = c.DB.Create(&category).Error

	if dbErr != nil {
		return dbErr
	}

	return nil
}

func (c *categoryDbStruct) DeleteCategory(id uuid.UUID) (errorMsg error) {
	defer func() {
		if r := recover(); r != nil {
			errorMsg = r.(error)
		}
	}()

	var dbCat domain.Category

	dbErr := c.DB.Model(&domain.Category{}).Where("id = ?", id).Find(&dbCat).Error

	if dbErr != nil {
		return dbErr
	}

	if dbCat.Name == "" {
		return fmt.Errorf("category does not exist")
	}

	dbErr = c.DB.Delete(dbCat).Error

	if dbErr != nil {
		return dbErr
	}

	return nil
}

func (c *categoryDbStruct) DeleteSubCategory(id uuid.UUID) (errorMsg error) {
	defer func() {
		if r := recover(); r != nil {
			errorMsg = r.(error)
		}
	}()

	var dbCat domain.SubCategory

	dbErr := c.DB.Model(&domain.SubCategory{}).Where("id = ?", id).Find(&dbCat).Error

	if dbErr != nil {
		return dbErr
	}

	if dbCat.Name == "" {
		return fmt.Errorf("sub category does not exist")
	}

	dbErr = c.DB.Delete(&dbCat).Error

	if dbErr != nil {
		return dbErr
	}

	return nil
}

func (c *categoryDbStruct) ListCategory() (responseCats []domain.CategoryResponse, errorMsg error) {
	defer func() {
		if r := recover(); r != nil {
			errorMsg = r.(error)
		}

	}()

	var categories []domain.Category

	dbErr := c.DB.Find(&categories).Error

	if dbErr != nil {
		return responseCats, dbErr
	}

	for _, category := range categories {

		subCats := []domain.SubCategory{}

		dbErr := c.DB.Model(&domain.SubCategory{}).Where("category_id = ?", category.ID).Find(&subCats).Error

		if dbErr != nil {
			return responseCats, dbErr
		}

		categoryResponse := category.ToResponse()

		categoryResponse.SubCategories = []domain.SubCategoryResponse{}

		for _, subCat := range subCats {
			categoryResponse.SubCategories = append(categoryResponse.SubCategories, subCat.ToResponse())
		}

		responseCats = append(responseCats, categoryResponse)

	}

	return responseCats, nil

}

func (c *categoryDbStruct) UpdateCategory(category *domain.Category) (errorMsg error) {
	defer func() {
		if r := recover(); r != nil {
			errorMsg = r.(error)
		}
	}()

	var dbCat domain.Category

	dbErr := c.DB.Model(&domain.Category{}).Where("id = ?", category.ID).Find(&dbCat).Error

	if dbErr != nil {
		return dbErr
	}

	if dbCat.Name == "" {
		return fmt.Errorf("category does not exist")

	}

	dbErr = c.DB.Updates(&category).Error

	if dbErr != nil {
		return dbErr
	}

	return nil

}

func (c *categoryDbStruct) UpdateSubCategory(category *domain.SubCategory) (errorMsg error) {
	defer func() {
		if r := recover(); r != nil {
			errorMsg = r.(error)
		}
	}()

	var dbCat domain.SubCategory

	dbErr := c.DB.Model(&domain.SubCategory{}).Where("id = ?", category.ID).Find(&dbCat).Error

	if dbErr != nil {
		return dbErr
	}

	if dbCat.Name == "" {
		return fmt.Errorf("sub category does not exist")

	}

	dbErr = c.DB.Updates(&category).Error

	if dbErr != nil {
		return dbErr
	}

	return nil
}
