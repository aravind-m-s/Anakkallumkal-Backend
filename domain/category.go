package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name string    `gorm:"not null;default:''"`
}

type CategoryResponse struct {
	ID            uuid.UUID             `json:"id"`
	Name          string                `json:"name"`
	SubCategories []SubCategoryResponse `json:"sub_categories"`
}

func (c *Category) ToResponse() CategoryResponse {
	return CategoryResponse{
		ID:   c.ID,
		Name: c.Name,
	}
}
