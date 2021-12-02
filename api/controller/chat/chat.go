package chat

import (
	"messenger-app/api/middlewares"
	"messenger-app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ChatController struct {
	chatModel         models.ChatModel
	userModel         models.UserModel
	conversationModel models.ConversationModel
}

func NewChatController(chatModel models.ChatModel, userModel models.UserModel, conversationModel models.ConversationModel) *ChatController {
	return &ChatController{
		chatModel,
		userModel,
		conversationModel,
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

	user1, _ := controller.userModel.GetNameById(int(userId))
	user2, _ := controller.userModel.GetNameById(int(chat.ReceiverID))

	conversation := models.Conversation{
		Name:        user1 + "-" + user2,
		UnreadCount: +1,
		UserID:      int(userId),
		ChatID:      int(data.ID),
	}

	_, err = controller.conversationModel.CreateConversation(conversation)
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

func (controller *ChatController) GetConversation(c echo.Context) error {
	userId := middlewares.ExtractTokenUser(c)

	data, err := controller.chatModel.GetConversation(int(userId))

	var dataNameReceiver, dataConv []string
	for i := range data {
		dataNameTemp, _ := controller.userModel.GetNameById(data[i])
		dataNameReceiver = append(dataNameReceiver, dataNameTemp)
	}
	nameUser, _ := controller.userModel.GetNameById(int(userId))

	for _, v := range dataNameReceiver {
		vConcate := nameUser + "-" + v
		dataConv = append(dataConv, vConcate)
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
		"data":    dataConv,
	})

}

func (controller *ChatController) GetLastMessage(c echo.Context) error {
	userId := middlewares.ExtractTokenUser(c)

	receiverId, err := controller.chatModel.GetConversation(int(userId))
	var messageData []string
	for i := range receiverId {
		data, _ := controller.chatModel.GetLastMessage(int(userId), receiverId[i])
		messageData = append(messageData, data)
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
		"data":    messageData,
	})

}
