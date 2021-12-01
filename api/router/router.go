package router

import (
	"messenger-app/api/controller/users"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo, userController *users.UserController) {
	// Authorization
	e.POST("/register", userController.RegisterUserController)
	e.POST("/login", userController.LoginUserController)
}
