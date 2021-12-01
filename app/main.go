package main

import (
	"fmt"
	"messenger-app/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := config.GetConfig()

	// db := util.MysqlDatabaseConnection

	e := echo.New()
	// middleware.LoggerMiddleware(e)

	address := fmt.Sprintf("localhost:%d", config.Port)

	if err := e.Start(address); err != nil {
		log.Info("shutting down the server")
	}
}
