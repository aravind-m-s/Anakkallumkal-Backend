package handlers

import (
	"github.com/aravind-m-s/anakallumkal-backend/config"
	"github.com/aravind-m-s/anakallumkal-backend/service"
	"github.com/gin-gonic/gin"
)

type FurnitureHandlerStruct struct {
	service service.FurnitureServiceInterface
	cnf     *config.EnvModel
}

func InitFurnitureHandler(service service.FurnitureServiceInterface, cnf *config.EnvModel) *FurnitureHandlerStruct {
	return &FurnitureHandlerStruct{service: service, cnf: cnf}
}

func (a *FurnitureHandlerStruct) ListFurniture(c *gin.Context) {
	a.service.ListFurniture(c)
}

func (a *FurnitureHandlerStruct) CreateFurniture(c *gin.Context) {
	a.service.CreateFurniture(c)
}

func (a *FurnitureHandlerStruct) UpdateFurniture(c *gin.Context) {
	a.service.UpdateFurniture(c)
}

func (a *FurnitureHandlerStruct) DeleteFurniture(c *gin.Context) {
	a.service.DeleteFurniture(c)
}

func (a *FurnitureHandlerStruct) ExportFurnitures(c *gin.Context) {
	a.service.ExportFurnitures(c)
}
