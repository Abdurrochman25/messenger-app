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
	userModel models.UserModel
}

type ConversationResponse struct {
	Conversation string `json:"conversation" form:"conversation"`
	Message      string `json:"message" form:"message"`
	Unread       int    `json:"unread" form:"unread"`
}

func NewChatController(chatModel models.ChatModel, userModel models.UserModel) *ChatController {
	return &ChatController{
		chatModel,
		userModel,
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

	_, err = controller.chatModel.UpdateUnreadMessage(int(userId), receiverId)
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

func (controller *ChatController) GetConversation(c echo.Context) error {
	userId := middlewares.ExtractTokenUser(c)
	// Get User Name
	nameUser, _ := controller.userModel.GetNameById(int(userId))

	// Get Receiver Id
	receiverId, err := controller.chatModel.GetConversation(int(userId))

	var dataMap ConversationResponse
	var data []interface{}

	for i := range receiverId {
		dataNameTemp, _ := controller.userModel.GetNameById(receiverId[i])
		vConcate := nameUser + "-" + dataNameTemp
		message, _ := controller.chatModel.GetLastMessage(int(userId), receiverId[i])
		counter, _ := controller.chatModel.GetCountUnreadMessage(int(userId), receiverId[i])

		dataMap = ConversationResponse{
			Conversation: vConcate,
			Message:      message,
			Unread:       counter,
		}
		data = append(data, dataMap)
	}

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
