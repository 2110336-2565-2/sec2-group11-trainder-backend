package controllers

import (
	"fmt"
	"net/http"
	"trainder-api/models"
	"trainder-api/responses"
	"trainder-api/utils/tokens"

	"github.com/gin-gonic/gin"
)

// @Summary Get roomID to communicate with audience
// @Description Get roomID to communicate with audience (can omit this function by using the roomID format trainer_{trainerUsername}_trainee_{traineeUsername})  NOTICE THAT all time in chat is at UTC
// @Tags chats
// @Accept json
// @Produce json
// @Param audience query string true "audience of this conversation (username)"
// @Security BearerAuth
// @Success 200 {object} responses.ChatRoomIDResponse
// @Failure 400 {object} responses.ChatRoomIDResponse
// @Router /protected/get-RoomID [GET]
func GetRoomID() gin.HandlerFunc {
	return func(c *gin.Context) {
		audience := c.Query("audience")
		username, err := tokens.ExtractTokenUsername(c)
		var rid string
		if models.IsTrainer(username) {
			rid = fmt.Sprintf("trainer_%s_trainee_%s", username, audience)
		} else {
			rid = fmt.Sprintf("trainer_%s_trainee_%s", audience, username)
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ChatRoomIDResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
		}

		c.JSON(http.StatusOK, responses.ChatRoomIDResponse{
			Status: http.StatusOK,
			RoomID: rid,
		})
	}
}

// @Summary Get all chat on sidebar that user communicate with
// @Description Get all chat on sidebar that user communicate with with their latest message NOTICE THAT all time in chat is at UTC
// @Tags chats
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.AllChatResponse
// @Failure 400 {object} responses.AllChatResponse
// @Router /protected/get-All-Chats [GET]
func GetChatsAndLatestMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.AllChatResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
		}
		result, err := models.GetAllChatLatestMessage(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AllChatResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, responses.AllChatResponse{
			Status:  http.StatusOK,
			Message: `success!`,
			AllChat: result,
		})
	}
}

// @Summary Get all messages that user communicate with specific audience
// @Description Get all messages that user communicate with specific audience NOTICE THAT all time in chat is at UTC
// @Tags chats
// @Accept json
// @Produce json
// @Param audience query string true "audience of this conversation (username)"
// @Security BearerAuth
// @Success 200 {object} responses.PastChatResponse
// @Failure 400 {object} responses.PastChatResponse
// @Router /protected/get-past-chat [GET]
func GetPastChat() gin.HandlerFunc {
	return func(c *gin.Context) {
		audience := c.Query("audience")
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.PastChatResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
		}
		result, err := models.GetPastChat(username, audience)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PastChatResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, responses.PastChatResponse{
			Status:       http.StatusOK,
			Message:      `success!`,
			ChatMessages: result,
		})
	}
}
