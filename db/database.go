package database

import (
	"fmt"

	"github.com/aravind-m-s/anakallumkal-backend/config"
	"github.com/aravind-m-s/anakallumkal-backend/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(cnf *config.EnvModel) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cnf.DbHost, cnf.DbUser, cnf.DbName, cnf.DbPort, cnf.DbPassword)
	db, err := gorm.Open(postgres.Open(psqlInfo))

	db.AutoMigrate(&domain.Brand{})
	db.AutoMigrate(&domain.Furniture{})
	db.AutoMigrate(&domain.Shop{})
	db.AutoMigrate(&domain.Category{})

	return db, err

}
