package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name   string    `gorm:"not null;default:''"`
	Image  string    `gorm:"not null;default:''"`
	ShopID uuid.UUID `gorm:"not null;index"`
	Shop   Shop      `gorm:"foreignKey:ShopID"`
	Count  int64
}

type BrandResponse struct {
	ID    uuid.UUID    `json:"id"`
	Name  string       `json:"name"`
	Image string       `json:"image"`
	Shop  ShopResponse `json:"shop"`
	Count int64        `json:"count"`
}

func (b *Brand) ToResponse() BrandResponse {
	return BrandResponse{
		Name:  b.Name,
		ID:    b.ID,
		Shop:  b.Shop.ToResponse(),
		Image: b.Image,
		Count: b.Count,
	}
}

