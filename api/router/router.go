package router

import (
	"messenger-app/api/controller/chat"
	"messenger-app/api/controller/users"
	"messenger-app/constant"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, userController *users.UserController, chatController *chat.ChatController) {
	// Authorization
	e.POST("/register", userController.RegisterUserController)
	e.POST("/login", userController.LoginUserController)

	// JWT Middleware
	jwtMiddleware := middleware.JWT([]byte(constant.SECRET_JWT))

	e.POST("/chat", chatController.SendMessageController, jwtMiddleware)
	e.GET("/chat/:receiverId", chatController.GetMessageByReceiverId, jwtMiddleware)
	e.GET("/chat", chatController.GetAllMessage, jwtMiddleware)
	e.GET("/conversation", chatController.GetConversation, jwtMiddleware)
}
