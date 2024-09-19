package handlers

import (
	"github.com/aravind-m-s/anakallumkal-backend/service"
	"github.com/gin-gonic/gin"
)

type CategoryHandlerStruct struct {
	service service.CategoryService
}

func InitCategoryHandler(service service.CategoryService) *CategoryHandlerStruct {
	return &CategoryHandlerStruct{service: service}
}

func (a *CategoryHandlerStruct) ListCategory(c *gin.Context) {
	a.service.ListCategory(c)
}

func (a *CategoryHandlerStruct) CreateCategory(c *gin.Context) {
	a.service.CreateCategory(c)
}

func (a *CategoryHandlerStruct) CreateSubCategory(c *gin.Context) {
	a.service.CreateSubCategory(c)
}

func (a *CategoryHandlerStruct) UpdateCategory(c *gin.Context) {
	a.service.UpdateCategory(c)
}

func (a *CategoryHandlerStruct) UpdateSubCategory(c *gin.Context) {
	a.service.UpdateSubCategory(c)
}

func (a *CategoryHandlerStruct) DeleteCategory(c *gin.Context) {
	a.service.DeleteCategory(c)
}

func (a *CategoryHandlerStruct) DeleteSubCategory(c *gin.Context) {
	a.service.DeleteSubCategory(c)
}
