package service

import (
	"fmt"
	"net/http"

	"github.com/aravind-m-s/anakallumkal-backend/domain"
	"github.com/aravind-m-s/anakallumkal-backend/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryService interface {
	CreateCategory(c *gin.Context)
	CreateSubCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	UpdateSubCategory(c *gin.Context)
	DeleteSubCategory(c *gin.Context)
	ListCategory(c *gin.Context)
}

type categoryServiceStruct struct {
	repo repository.CategoryRepository
}

func InitCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryServiceStruct{repo: repo}
}

func (s *categoryServiceStruct) CreateCategory(c *gin.Context) {

	name := c.PostForm("name")

	if name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Name is required"})
		return
	}

	category := domain.Category{
		Name: name,
	}

	err := s.repo.CreateCategory(&category)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resCat := category.ToResponse()

	resCat.SubCategories = []domain.SubCategoryResponse{}

	c.JSON(http.StatusOK, resCat)
}

func (s *categoryServiceStruct) CreateSubCategory(c *gin.Context) {
	name := c.PostForm("name")
	idStr := c.PostForm("id")

	errorMap := gin.H{}

	if name == "" {
		errorMap["name"] = "Name is required"
	}

	if idStr == "" {
		errorMap["id"] = "Category is required"
	}

	id, err := uuid.Parse(idStr)

	if err != nil {
		errorMap["id"] = "Invalid Category"
	}

	if len(errorMap) != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorMap)
		return
	}

	category := domain.SubCategory{
		Name:       name,
		CategoryID: id,
	}

	err = s.repo.CreateSubCategory(&category)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category.ToResponse())
}

func (s *categoryServiceStruct) DeleteCategory(c *gin.Context) {
	idStr := c.PostForm("id")

	errorMap := gin.H{}

	if idStr == "" {
		errorMap["id"] = "Category is required"
	}

	id, err := uuid.Parse(idStr)

	if err != nil {
		errorMap["id"] = "Invalid Category"
	}

	if len(errorMap) != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorMap)
		return
	}

	err = s.repo.DeleteCategory(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func (s *categoryServiceStruct) DeleteSubCategory(c *gin.Context) {
	idStr := c.PostForm("id")

	errorMap := gin.H{}

	if idStr == "" {
		errorMap["id"] = "Sub Category is required"
	}

	id, err := uuid.Parse(idStr)

	if err != nil {
		errorMap["id"] = "Invalid Sub Category"
	}

	if len(errorMap) != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorMap)
		return
	}

	err = s.repo.DeleteSubCategory(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sub Category deleted successfully"})
}

func (s *categoryServiceStruct) ListCategory(c *gin.Context) {
	categories, err := s.repo.ListCategory()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)

}

func (s *categoryServiceStruct) UpdateCategory(c *gin.Context) {
	name := c.PostForm("name")
	idStr := c.PostForm("id")

	errorMap := gin.H{}

	if name == "" {
		errorMap["name"] = "Name is required"
	}

	if idStr == "" {
		errorMap["id"] = "Id is required"
	}

	id, err := uuid.Parse(idStr)

	if err != nil {
		errorMap["id"] = "Invalid Id"
	}

	if len(errorMap) != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorMap)
		return
	}

	category := domain.Category{
		Name: name,
		ID:   id,
	}

	err = s.repo.UpdateCategory(&category)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resCat := category.ToResponse()

	resCat.SubCategories = []domain.SubCategoryResponse{}

	c.JSON(http.StatusOK, resCat)
}

func (s *categoryServiceStruct) UpdateSubCategory(c *gin.Context) {
	name := c.PostForm("name")
	idStr := c.PostForm("id")
	catIdStr := c.PostForm("category_id")

	errorMap := gin.H{}

	if name == "" {
		errorMap["name"] = "Name is required"
	}

	if idStr == "" {
		errorMap["id"] = "Id is required"
	}

	if catIdStr == "" {
		errorMap["id"] = "Category Id is required"
	}

	id, err := uuid.Parse(idStr)

	if err != nil && idStr != "" {
		errorMap["id"] = "Invalid Id"
	}

	catId, err := uuid.Parse(catIdStr)

	if err != nil && catIdStr != "" {
		errorMap["id"] = "Invalid Category Id"
	}

	if len(errorMap) != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorMap)
		return
	}

	category := domain.SubCategory{
		Name:       name,
		ID:         id,
		CategoryID: catId,
	}

	fmt.Printf("category: %v\n", category)

	err = s.repo.UpdateSubCategory(&category)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category.ToResponse())
}
