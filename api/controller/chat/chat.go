package chat

import (
	"messenger-app/api/middlewares"
	"messenger-app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ChatController struct {
	chatModel models.ChatModel
}

func NewChatController(chatModel models.ChatModel) *ChatController {
	return &ChatController{
		chatModel,
	}
}

func (controller *ChatController) SendMessageController(c echo.Context) error {
	var chatRequest models.Chat

	if err := c.Bind(&chatRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	userId := middlewares.ExtractTokenUser(c)

	chat := models.Chat{
		ReceiverID: chatRequest.ReceiverID,
		Message:    chatRequest.Message,
		UserID:     int(userId),
	}

	data, err := controller.chatModel.SendMessage(chat)
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
		"message": "Success Send Message",
		"data":    data.Message,
	})
}

func (controller *ChatController) GetMessageByReceiverId(c echo.Context) error {
	userId := middlewares.ExtractTokenUser(c)

	receiverId, err := strconv.Atoi(c.Param("receiverId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	data, err := controller.chatModel.GetMessageByReceiverId(int(userId), receiverId)
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
		"message": "Success Send Message",
		"data":    data,
	})

}

func (controller *ChatController) GetAllMessage(c echo.Context) error {
	userId := middlewares.ExtractTokenUser(c)

	data, err := controller.chatModel.GetAllMessage(int(userId))
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
		"message": "Success Send Message",
		"data":    data,
	})

}
