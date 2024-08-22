package handlers

import (
	"net/http"

	"github.com/aravind-m-s/anakallumkal-backend/service"
	"github.com/gin-gonic/gin"
)

type SeederHandlerStruct struct {
	service service.SeederService
}

func InitSeederHandler(service service.SeederService) *SeederHandlerStruct {
	return &SeederHandlerStruct{service: service}
}

func (a *SeederHandlerStruct) ShopSeeder(c *gin.Context) {
	value := a.service.ShopSeeder()

	statusCode := http.StatusOK

	if value != "Success" {
		statusCode = http.StatusBadRequest
	}

	c.JSON(statusCode, gin.H{"message": value})

}

func (a *SeederHandlerStruct) ShopGet(c *gin.Context) {
	value, err := a.service.ShopGet()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, value)
	}

}
