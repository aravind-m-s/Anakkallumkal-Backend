package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name  string    `gorm:"not null;default:''"`
	Place string    `gorm:"not null;default:''"`
	Image string    `gorm:"not null;default:''"`
}

type ShopResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Place string    `json:"place"`
	Image string    `json:"image"`
}

func (s *Shop) ToResponse() ShopResponse {
	return ShopResponse{
		ID:    s.ID,
		Image: s.Image,
		Place: s.Place,
		Name:  s.Name,
	}
}
