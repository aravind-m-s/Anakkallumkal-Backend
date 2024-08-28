package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubCategory struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name       string    `gorm:"not null;default:''"`
	CategoryID uuid.UUID `gorm:"not null;index"`
	Category   Category  `gorm:"foreignKey:CategoryID"`
}

type SubCategoryResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Category Category  `json:"category"`
}

func (c *SubCategory) ToResponse() SubCategoryResponse {
	return SubCategoryResponse{
		ID:       c.ID,
		Name:     c.Name,
		Category: c.Category,
	}
}
