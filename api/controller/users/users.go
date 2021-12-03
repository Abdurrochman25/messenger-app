package users

import (
	"messenger-app/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userModel models.UserModel
}

func NewUserController(userModel models.UserModel) *UserController {
	return &UserController{
		userModel,
	}
}

func (controller *UserController) RegisterUserController(c echo.Context) error {
	var userRequest models.User

	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	_, err := controller.userModel.Register(userRequest)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"code":    500,
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"code":    200,
		"message": "Success Register Account",
	})
}

func (controller *UserController) LoginUserController(c echo.Context) error {
	var userRequest models.User

	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	user, err := controller.userModel.Login(userRequest.Username, userRequest.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"code":    500,
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"code":    200,
		"message": "Success Login Account",
		"token":   user.Token,
	})
}
