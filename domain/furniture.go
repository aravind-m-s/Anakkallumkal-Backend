package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Furniture struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"not null;default:''"`
	Image     string    `gorm:"not null;default:''"`
	ProductNo string    `gorm:"not null;default:''"`
	BrandID   uuid.UUID `gorm:"not null;index"`
	Brand     Brand     `gorm:"foreignKey:BrandID"`
	ShopID    uuid.UUID `gorm:"not null;index"`
	Shop      Shop      `gorm:"foreignKey:ShopID"`
	Stock     int       `gorm:"not null;default:0"`
	Price     int       `gorm:"not null;default:0"`
}
