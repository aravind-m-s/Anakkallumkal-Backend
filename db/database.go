package database

import (
	"github.com/aravind-m-s/anakallumkal-backend/config"
	"github.com/aravind-m-s/anakallumkal-backend/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(cnf *config.EnvModel) (*gorm.DB, error) {
	// psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cnf.DbHost, cnf.DbUser, cnf.DbName, cnf.DbPort, cnf.DbPassword)
	psqlInfo := cnf.DBUrl
	db, err := gorm.Open(postgres.Open(psqlInfo))

	if err != nil {
		return db, err
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	db.AutoMigrate(&domain.Brand{})
	db.AutoMigrate(&domain.Furniture{})
	db.AutoMigrate(&domain.Shop{})
	db.AutoMigrate(&domain.Category{})

	return db, err

}
