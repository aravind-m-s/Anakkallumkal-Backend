package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Furniture struct {
	gorm.Model
	ID            uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name          string      `gorm:"not null;default:''"`
	Image         string      `gorm:"not null;default:''"`
	ProductNo     string      `gorm:"not null;default:''"`
	BrandID       uuid.UUID   `gorm:"not null;index"`
	Brand         Brand       `gorm:"foreignKey:BrandID"`
	SubCategoryID uuid.UUID   `gorm:"not null;index"`
	SubCategory   SubCategory `gorm:"foreignKey:SubCategoryID"`
	Stock         int         `gorm:"not null;default:0"`
	Price         int         `gorm:"not null;default:0"`
	Rows          int         `gorm:"not null;default:0"`
}

type FurnitureResponse struct {
	Name      string              `json:"name"`
	ID        uuid.UUID           `json:"id"`
	Image     string              `json:"image"`
	ProductNo string              `json:"product_no"`
	Brand     BrandResponse       `json:"brand"`
	Category  SubCategoryResponse `json:"category"`
	Stock     int                 `json:"stock"`
	Price     int                 `json:"price"`
	Rows      int                 `json:"rows"`
}

func (f *Furniture) ToResponse() FurnitureResponse {
	return FurnitureResponse{
		ID:        f.ID,
		Image:     f.Image,
		Name:      f.Name,
		ProductNo: f.ProductNo,
		Brand:     f.Brand.ToResponse(),
		Category:  f.SubCategory.ToResponse(),
		Stock:     f.Stock,
		Price:     f.Price,
		Rows:      f.Rows,
	}
}
