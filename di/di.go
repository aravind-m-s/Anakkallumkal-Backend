package di

import (
	"github.com/aravind-m-s/anakallumkal-backend/api"
	"github.com/aravind-m-s/anakallumkal-backend/api/middlewares"
	"github.com/aravind-m-s/anakallumkal-backend/common"
	"github.com/aravind-m-s/anakallumkal-backend/config"
)

func InitServer(cnf *config.EnvModel) (*api.ServerHTTP, error) {
	// db, err := database.InitDatabase(cnf)

	jwt := common.NewHelper(cnf)
	authorization := middlewares.NewAuthorization(jwt)

	server := api.Handler(authorization)

	// if err != nil {
	// 	return nil, err
	// }

	return server, nil
}
