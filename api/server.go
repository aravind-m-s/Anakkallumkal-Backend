package api

import (
	"log"
	"net/http"

	"github.com/aravind-m-s/anakallumkal-backend/common"
	"github.com/aravind-m-s/anakallumkal-backend/config"
	database "github.com/aravind-m-s/anakallumkal-backend/db"
	"github.com/aravind-m-s/anakallumkal-backend/repository"
	"github.com/aravind-m-s/anakallumkal-backend/server/handlers"
	"github.com/aravind-m-s/anakallumkal-backend/server/middlewares"
	"github.com/aravind-m-s/anakallumkal-backend/service"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewHandler(w http.ResponseWriter, r *http.Request) {

	config := config.InitConfig()

	db, err := database.InitDatabase(config)

	if err != nil {
		panic(err)
	}

	jwt := common.NewHelper(config)
	authorization := middlewares.NewAuthorization(jwt)

	furnitureRepo := repository.InitFurnitureRepo(db)
	furnitureService := service.InitFurnitureService(furnitureRepo)
	furnitureHandler := handlers.InitFurnitureHandler(furnitureService, config)

	brandRepo := repository.InitBrandRepo(db)
	brandService := service.InitBrandService(brandRepo)
	brandHandler := handlers.InitBrandHandler(brandService, config)

	seederRepo := repository.InitSeederRepo(db)
	seederService := service.InitSeederService(seederRepo)
	seederHandler := handlers.InitSeederHandler(seederService)

	server := Handler(furnitureHandler, brandHandler, seederHandler, authorization)

	if err != nil {
		log.Fatal("Uanble to connect to db", err)
	} else {
		server.Start(config)
	}
}

func Handler(furnitureHandler *handlers.FurnitureHandlerStruct, brandHandler *handlers.BrandHandlerStruct, seederHandler *handlers.SeederHandlerStruct, middlewares *middlewares.AuthorizationStruct) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())
	engine.Static("/media", "./media")

	apiGroup := engine.Group("/api")

	furnitureGroup := apiGroup.Group("/furniture")

	furnitureGroup.GET("/list/:id", furnitureHandler.ListFurniture)
	furnitureGroup.POST("/create", furnitureHandler.CreateFurniture)
	furnitureGroup.PUT("/:id", furnitureHandler.UpdateFurniture)
	furnitureGroup.DELETE("/:id", furnitureHandler.DeleteFurniture)
	furnitureGroup.GET("/export/:id", furnitureHandler.ExportFurnitures)

	brandGroup := apiGroup.Group("/brand")

	brandGroup.GET("/list/", brandHandler.ListBrand)
	brandGroup.POST("/create", brandHandler.CreateBrand)
	brandGroup.PUT("/:id", brandHandler.UpdateBrand)
	brandGroup.DELETE("/:id", brandHandler.DeleteBrand)

	seederGroup := apiGroup.Group("/seeder")

	seederGroup.GET("/shop", seederHandler.ShopSeeder)
	seederGroup.GET("/shop/list", seederHandler.ShopGet)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start(cnf *config.EnvModel) {
	sh.engine.Run(cnf.Port)
}
