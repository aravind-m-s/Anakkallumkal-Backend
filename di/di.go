package di

import (
	"github.com/aravind-m-s/anakallumkal-backend/api"
	"github.com/aravind-m-s/anakallumkal-backend/api/handlers"
	"github.com/aravind-m-s/anakallumkal-backend/api/middlewares"
	"github.com/aravind-m-s/anakallumkal-backend/common"
	"github.com/aravind-m-s/anakallumkal-backend/config"
	database "github.com/aravind-m-s/anakallumkal-backend/db"
	"github.com/aravind-m-s/anakallumkal-backend/repository"
	"github.com/aravind-m-s/anakallumkal-backend/service"
)

func InitServer(cnf *config.EnvModel) (*api.ServerHTTP, error) {
	db, err := database.InitDatabase(cnf)

	if err != nil {
		return nil, err
	}

	jwt := common.NewHelper(cnf)
	authorization := middlewares.NewAuthorization(jwt)

	furnitureRepo := repository.InitFurnitureRepo(db)
	furnitureService := service.InitFurnitureService(furnitureRepo)
	furnitureHandler := handlers.InitFurnitureHandler(furnitureService, cnf)

	brandRepo := repository.InitBrandRepo(db)
	brandService := service.InitBrandService(brandRepo)
	brandHandler := handlers.InitBrandHandler(brandService, cnf)

	seederRepo := repository.InitSeederRepo(db)
	seederService := service.InitSeederService(seederRepo)
	seederHandler := handlers.InitSeederHandler(seederService)

	server := api.Handler(furnitureHandler, brandHandler, seederHandler, authorization)

	return server, nil
}
