package service

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	_ "image/jpeg"
	_ "image/png"

	"github.com/aravind-m-s/anakallumkal-backend/domain"
	"github.com/aravind-m-s/anakallumkal-backend/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

type FurnitureServiceInterface interface {
	CreateFurniture(c *gin.Context)
	DeleteFurniture(c *gin.Context)
	ListFurniture(c *gin.Context)
	UpdateFurniture(c *gin.Context)
	ExportFurnitures(c *gin.Context)
}

type furnitureServiceStruct struct {
	repo repository.FurnitureRepoInterface
}

func InitFurnitureService(repo repository.FurnitureRepoInterface) FurnitureServiceInterface {
	return &furnitureServiceStruct{repo: repo}
}

func (f *furnitureServiceStruct) CreateFurniture(c *gin.Context) {

	defer func() {
		if r := recover(); r != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}
	}()

	name := c.PostForm("name")
	productNo := c.PostForm("product_no")
	stock := c.PostForm("stock")
	price := c.PostForm("price")
	rows := c.PostForm("rows")
	brand := c.PostForm("brand")
	category := c.PostForm("category")

	imageData, err := c.FormFile("image")

	errorMap := gin.H{}

	if name == "" {
		errorMap["name"] = "Name Cannot be empty"
	}

	if productNo == "" {
		errorMap["product_no"] = "Product Number Cannot be empty"
	}

	if stock == "" {
		stock = "0"
	}

	stockAmount, stockErr := strconv.Atoi(stock)

	if stockErr != nil {
		errorMap["stock"] = "Invalid stock amount"
	}

	if price == "" {
		price = "0"
	}

	priceAmount, priceErr := strconv.Atoi(price)

	if priceErr != nil {
		errorMap["price"] = "Invalid price amount"
	}

	if rows == "" {
		rows = "1"
	}

	rowsAmount, rowsErr := strconv.Atoi(rows)

	if rowsErr != nil {
		errorMap["rows"] = "Invalid rows amount"
	}

	if brand == "" {
		errorMap["brand"] = "Brand Cannot be empty"
	}

	brandId, brandErr := uuid.Parse(brand)

	if brandErr != nil {
		errorMap["brand"] = "Invalid brand id"
	}

	if category == "" {
		errorMap["category"] = "Category Cannot be empty"
	}

	categoryId, categoryErr := uuid.Parse(category)

	if categoryErr != nil {
		errorMap["category"] = "Invalid category id"
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

	furniture, err := f.repo.CreateFurniture(name, imagePath, productNo, brandId, stockAmount, priceAmount, categoryId, rowsAmount)

	if err != nil {
		os.Remove(imagePath)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, furniture.ToResponse())

}

func (f *furnitureServiceStruct) DeleteFurniture(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}
	}()

	id := c.Param("id")

	errorMap := gin.H{}

	if id == "" {
		errorMap["id"] = "Furniture ID is required"
	}

	furnitureId, furnitureErr := uuid.Parse(id)

	if furnitureErr != nil {
		errorMap["furniture"] = "Invalid furniture id"
	}

	if len(errorMap) != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorMap)
		return
	}

	err := f.repo.DeleteFurniture(furnitureId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Furniture deleted successfully"})
	}
}

func (f *furnitureServiceStruct) ExportFurnitures(c *gin.Context) {
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
		errorMap["furniture"] = "Invalid brand id"
	}

	if len(errorMap) != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorMap)
		return
	}

	furnitures, err := f.repo.ExportFurniture(brandId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	excelBrochure, err := convertToExcelBrochure(furnitures)
	excel, err := convertToExcelStockBook(furnitures)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"brochure_excel": c.Request.Host + "/media/" + excelBrochure, "stock_book": c.Request.Host + "/media/" + excel})
}
func convertToExcelBrochure(furnitures []domain.Furniture) (string, error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheetName := "Sheet1"

	f.NewSheet(sheetName)

	cmToExcelWidth := func(cm float64) float64 {
		return cm * 3.93700787 // 1 cm = 1 / 2.54 inch
	}

	cmToExcelHeight := func(cm float64) float64 {
		return cm * 25.7
	}

	colWidths := map[string]float64{
		"A": 5.20,
		"B": 8.15,
		"C": 3.85,
	}

	for col, width := range colWidths {
		if err := f.SetColWidth(sheetName, col, col, cmToExcelWidth(width)); err != nil {
			fmt.Println(err)
		}
	}

	borderStyleOdd := excelize.Style{
		Alignment: &excelize.Alignment{
			Vertical:   "center",
			Horizontal: "center",
			WrapText:   true,
		},
		Border: []excelize.Border{
			{Type: "top", Color: "000000", Style: 1},
			{Type: "left", Color: "000000", Style: 1},
		},
	}

	borderStyleEven := excelize.Style{
		Alignment: &excelize.Alignment{
			Vertical:   "center",
			Horizontal: "center",
			WrapText:   true,
		},
		Border: []excelize.Border{
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "left", Color: "000000", Style: 1},
		},
	}

	styleIDOdd, err := f.NewStyle(&borderStyleOdd)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	styleIDEven, err := f.NewStyle(&borderStyleEven)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if err := f.SetCellStyle(sheetName, "A1", fmt.Sprintf("C%d", len(furnitures)+1), styleIDOdd); err != nil {
		fmt.Println(err)
		return "", err
	}

	f.MergeCell(sheetName, "A1", "C1")
	f.SetCellValue(sheetName, "A1", furnitures[0].Brand.Name)

	for i := 0; i < len(furnitures)*2; i++ {
		row := i + 2
		height := 4.0
		if i%2 != 0 {
			height = 1.0
		}
		if err := f.SetRowHeight(sheetName, row, cmToExcelHeight(height)); err != nil {
			fmt.Println(err)
		}
	}

	index := 0

	for i := 0; i < len(furnitures)*2; i++ {
		row := i + 2

		// For odd rows (Name row)
		if row%2 != 0 {
			if err := f.SetCellStyle(sheetName, fmt.Sprintf("A%d", row), fmt.Sprintf("A%d", row), styleIDOdd); err != nil {
				fmt.Println(err)
				return "", err
			}
			if err := f.SetCellStyle(sheetName, fmt.Sprintf("B%d", row), fmt.Sprintf("B%d", row), styleIDOdd); err != nil {
				fmt.Println(err)
				return "", err
			}
			f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), furnitures[index].Name)
			index++
		} else {
			if err := f.SetCellStyle(sheetName, fmt.Sprintf("A%d", row), fmt.Sprintf("A%d", row), styleIDEven); err != nil {
				fmt.Println(err)
				return "", err
			}
			if err := f.SetCellStyle(sheetName, fmt.Sprintf("B%d", row), fmt.Sprintf("B%d", row), styleIDEven); err != nil {
				fmt.Println(err)
				return "", err
			}
			if err := f.AddPicture(sheetName, fmt.Sprintf("A%d", row), furnitures[index].Image, &excelize.GraphicOptions{
				AutoFit: true,
			}); err != nil {
				fmt.Println(err.Error())
				return "", err
			}
		}
	}

	fileName := furnitures[0].Brand.Name + " " + "Brochure" + ".xlsx"

	if err := f.SaveAs("./media/" + fileName); err != nil {
		fmt.Println(err)
		return "", err
	}

	return fileName, nil
}

func convertToExcelStockBook(furnitures []domain.Furniture) (string, error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheetName := "Sheet1"

	f.NewSheet(sheetName)

	cmToExcelWidth := func(cm float64) float64 {
		return cm * 3.93700787 // 1 cm = 1 / 2.54 inch * 10 (approx)
	}

	cmToExcelHeight := func(cm float64) float64 {
		return cm * 25.7
	}

	colWidths := map[string]float64{
		"A": 4.20,
		"B": 3.85,
		"C": 4.15,
	}

	for col, width := range colWidths {
		if err := f.SetColWidth(sheetName, col, col, cmToExcelWidth(width)); err != nil {
			fmt.Println(err)
		}
	}

	borderStyle := excelize.Style{
		Alignment: &excelize.Alignment{
			Vertical:   "center",
			Horizontal: "center",
			WrapText:   true,
		},

		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	}

	borderStyleWithBold := excelize.Style{
		Alignment: &excelize.Alignment{
			Vertical:   "center",
			Horizontal: "center",
			WrapText:   true,
		},
		Font: &excelize.Font{
			Bold: true,
			Size: 18,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	}

	styleID, err := f.NewStyle(&borderStyle)
	boldStyleID, err := f.NewStyle(&borderStyleWithBold)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if err := f.SetRowHeight(sheetName, 1, cmToExcelHeight(1.5)); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle(sheetName, "A1", fmt.Sprintf("C%d", len(furnitures)+1), styleID); err != nil {
		fmt.Println(err)
		return "", err
	}

	if err := f.SetCellStyle(sheetName, "A1", "C1", boldStyleID); err != nil {
		fmt.Println(err)
		return "", err
	}

	f.MergeCell(sheetName, "A1", "C1")
	f.SetCellValue(sheetName, "A1", furnitures[0].Brand.Name)

	for i := 0; i < len(furnitures); i++ {
		if err := f.SetRowHeight(sheetName, i+2, cmToExcelHeight(2)); err != nil {
			fmt.Println(err)
		}
	}

	for i := 0; i < len(furnitures); i++ {

		f.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), furnitures[i].Name)
		if furnitures[i].Price != 0 {

			f.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), furnitures[i].Price)
		}
	}

	fileName := furnitures[0].Brand.Name + " " + "Stock Book" + ".xlsx"

	if err := f.SaveAs("./media/" + fileName); err != nil {
		fmt.Println(err)
		return "", err
	}

	return fileName, nil
}

func (f *furnitureServiceStruct) ListFurniture(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}
	}()

	id := c.Param("id")
	query := c.Query("search")

	errorMap := gin.H{}

	brandId, brandErr := uuid.Parse(id)

	if brandErr != nil {
		errorMap["brand"] = "Invalid brand id"
	}

	if len(errorMap) != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorMap)
		return
	}

	dbFurnitures, err := f.repo.ListFurniture(brandId, query)

	furnitures := []domain.FurnitureResponse{}

	for _, furniture := range dbFurnitures {
		fmt.Printf("furniture: %v\n", furniture.DeletedAt)

		furnitures = append(furnitures, furniture.ToResponse())
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, furnitures)
	}
}

func (f *furnitureServiceStruct) UpdateFurniture(c *gin.Context) {

	defer func() {
		if r := recover(); r != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}
	}()

	id := c.Param("id")
	name := c.PostForm("name")
	productNo := c.PostForm("product_no")
	stock := c.PostForm("stock")
	price := c.PostForm("price")
	rows := c.PostForm("rows")
	brand := c.PostForm("brand")
	category := c.PostForm("category")

	imageData, err := c.FormFile("image")

	errorMap := gin.H{}

	if id == "" {
		errorMap["id"] = "Furniture ID is required"
	}

	furnitureId, furnitureErr := uuid.Parse(id)

	if furnitureErr != nil {
		errorMap["furniture"] = "Invalid furniture id"
	}

	if name == "" {
		errorMap["name"] = "Name Cannot be empty"
	}

	if productNo == "" {
		errorMap["product_no"] = "Product Number Cannot be empty"
	}

	if stock == "" {
		stock = "0"
	}

	stockAmount, stockErr := strconv.Atoi(stock)

	if stockErr != nil {
		errorMap["stock"] = "Invalid stock amount"
	}

	if price == "" {
		price = "0"
	}

	priceAmount, priceErr := strconv.Atoi(price)

	if priceErr != nil {
		errorMap["price"] = "Invalid price amount"
	}

	if rows == "" {
		rows = "1"
	}

	rowsAmount, rowsErr := strconv.Atoi(rows)

	if rowsErr != nil {
		errorMap["rows"] = "Invalid rows amount"
	}

	if brand == "" {
		errorMap["brand"] = "Brand Cannot be empty"
	}

	brandId, brandErr := uuid.Parse(brand)

	if brandErr != nil {
		errorMap["brand"] = "Invalid brand id"
	}

	if category == "" {
		errorMap["category"] = "Category Cannot be empty"
	}

	categoryId, categoryErr := uuid.Parse(category)

	if categoryErr != nil {
		errorMap["category"] = "Invalid category id"
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
			println(saveErr.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": saveErr.Error()})
			return
		}
	}

	furniture, err := f.repo.UpdateFurniture(furnitureId, name, imagePath, productNo, brandId, stockAmount, priceAmount, categoryId, rowsAmount)

	if err != nil {
		if imagePath != "" {
			os.Remove(imagePath)
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, furniture.ToResponse())
}
