package service

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aravind-m-s/anakallumkal-backend/domain"
	"github.com/aravind-m-s/anakallumkal-backend/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BrandServiceInterface interface {
	CreateBrand(c *gin.Context)
	DeleteBrand(c *gin.Context)
	ListBrand(c *gin.Context)
	UpdateBrand(c *gin.Context)
}

type brandServiceStruct struct {
	repo repository.BrandRepoInterface
}

func InitBrandService(repo repository.BrandRepoInterface) BrandServiceInterface {
	return &brandServiceStruct{repo: repo}
}

func (f *brandServiceStruct) CreateBrand(c *gin.Context) {

	defer func() {
		if r := recover(); r != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}
	}()

	name := c.PostForm("name")
	shop := c.PostForm("shop")

	imageData, err := c.FormFile("image")

	errorMap := gin.H{}

	if name == "" {
		errorMap["name"] = "Name Cannot be empty"
	}

	if shop == "" {
		errorMap["shop"] = "Shop Cannot be empty"
	}

	shopId, shopErr := uuid.Parse(shop)

	if shopErr != nil {
		errorMap["shop"] = "Invalid shop id"
	}

	if err != nil || imageData == nil {
		errorMap["image"] = "Image should be valid"
	}

	if len(errorMap) != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorMap)
		return
	}

	imagePath := "./media/" + strconv.FormatInt(time.Now().UnixMilli(), 10) + "." + strings.Split(imageData.Filename, ".")[len(strings.Split(imageData.Filename, "."))-1]

	saveErr := c.SaveUploadedFile(imageData, imagePath)

	if saveErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Unable to proccess the image"})
		return
	}

	brand, err := f.repo.CreateBrand(name, imagePath, shopId)

	if err != nil {
		os.Remove(imagePath)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, brand.ToResponse())

}

func (f *brandServiceStruct) DeleteBrand(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}
	}()

	id := c.Param("id")

	errorMap := gin.H{}

	if id == "" {
		errorMap["id"] = "Brand ID is required"
	}

	brandId, brandErr := uuid.Parse(id)

	if brandErr != nil {
		errorMap["brand"] = "Invalid brand id"
	}

	if len(errorMap) != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorMap)
		return
	}



	err := f.repo.DeleteBrand(brandId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Brand deleted successfully"})
	}
}

func (f *brandServiceStruct) ListBrand(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}
	}()

	dbBrands, dbShops, err := f.repo.ListBrand()

	brands := []domain.BrandResponse{}

	for _, brand := range dbBrands {
		brands = append(brands, brand.ToResponse())
	}
	shops := []domain.ShopResponse{}

	for _, shop := range dbShops {
		shops = append(shops, shop.ToResponse())
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {

		c.JSON(http.StatusOK, gin.H{"shops": shops, "brands": brands})
	}
}

func (f *brandServiceStruct) UpdateBrand(c *gin.Context) {

	defer func() {
		if r := recover(); r != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}
	}()

	id := c.Param("id")
	name := c.PostForm("name")
	shop := c.PostForm("shop")

	imageData, err := c.FormFile("image")

	errorMap := gin.H{}

	if id == "" {
		errorMap["id"] = "Brand ID is required"
	}

	brandId, brandErr := uuid.Parse(id)

	if brandErr != nil {
		errorMap["brand"] = "Invalid brand id"
	}

	if name == "" {
		errorMap["name"] = "Name Cannot be empty"
	}

	if shop == "" {
		errorMap["shop"] = "Shop Cannot be empty"
	}

	shopId, shopErr := uuid.Parse(shop)

	if shopErr != nil {
		errorMap["shop"] = "Invalid shop id"
	}

	if len(errorMap) != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorMap)
		return
	}

	imagePath := ""

	if imageData != nil && err == nil {
		imagePath = "./media/" + strconv.FormatInt(time.Now().UnixMilli(), 10) + "." + strings.Split(imageData.Filename, ".")[len(strings.Split(imageData.Filename, "."))-1]
		saveErr := c.SaveUploadedFile(imageData, imagePath)
		if saveErr != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": saveErr.Error()})
			return
		}
	}

	brand, err := f.repo.UpdateBrand(brandId, name, imagePath, shopId)

	if err != nil {
		if imagePath != "" {
			os.Remove(imagePath)

		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, brand.ToResponse())
}
