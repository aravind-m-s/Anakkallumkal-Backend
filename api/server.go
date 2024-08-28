package api

import (
	"github.com/aravind-m-s/anakallumkal-backend/api/handlers"
	"github.com/aravind-m-s/anakallumkal-backend/api/middlewares"
	"github.com/aravind-m-s/anakallumkal-backend/config"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
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
