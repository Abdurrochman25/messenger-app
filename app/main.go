package main

import (
	"fmt"
	"messenger-app/config"
	"messenger-app/models"
	"messenger-app/util"

	userController "messenger-app/api/controller/users"
	"messenger-app/api/middlewares"
	"messenger-app/api/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := config.GetConfig()

	db := util.MysqlDatabaseConnection(config)

	userModel := models.NewUserModel(db)

	newUserController := userController.NewUserController(userModel)

	e := echo.New()
	middlewares.LoggerMiddlewares(e)

	router.Route(e, newUserController)

	address := fmt.Sprintf("localhost:%d", config.Port)

	if err := e.Start(address); err != nil {
		log.Info("shutting down the server")
	}
}
