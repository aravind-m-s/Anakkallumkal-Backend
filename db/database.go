package database

import (
	"github.com/aravind-m-s/anakallumkal-backend/config"
	"github.com/aravind-m-s/anakallumkal-backend/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(cnf *config.EnvModel) (*gorm.DB, error) {
	psqlInfo := cnf.DBUrl
	db, err := gorm.Open(postgres.Open(psqlInfo))

	if err != nil {
		return db, err
	}

	db.AutoMigrate(&domain.Brand{})
	db.AutoMigrate(&domain.Furniture{})
	db.AutoMigrate(&domain.Shop{})
	db.AutoMigrate(&domain.Category{})

	return db, err

}
