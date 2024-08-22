package handlers

import (
	"github.com/aravind-m-s/anakallumkal-backend/config"
	"github.com/aravind-m-s/anakallumkal-backend/service"
	"github.com/gin-gonic/gin"
)

type BrandHandlerStruct struct {
	service service.BrandServiceInterface
	cnf     *config.EnvModel
}

func InitBrandHandler(service service.BrandServiceInterface, cnf *config.EnvModel) *BrandHandlerStruct {
	return &BrandHandlerStruct{service: service, cnf: cnf}
}

func (a *BrandHandlerStruct) ListBrand(c *gin.Context) {
	a.service.ListBrand(c)
}

func (a *BrandHandlerStruct) CreateBrand(c *gin.Context) {
	a.service.CreateBrand(c)
}

func (a *BrandHandlerStruct) UpdateBrand(c *gin.Context) {
	a.service.UpdateBrand(c)
}

func (a *BrandHandlerStruct) DeleteBrand(c *gin.Context) {
	a.service.DeleteBrand(c)
}
