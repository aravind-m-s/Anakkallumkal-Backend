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
}
