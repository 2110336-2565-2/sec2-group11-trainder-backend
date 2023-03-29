package controllers

import (
	"fmt"
	"net/http"
	"trainder-api/models"
	"trainder-api/responses"
	"trainder-api/utils/tokens"

	"github.com/gin-gonic/gin"
)

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

func GetChatsAndLatestMessege() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.AllChatResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
		}
		result, err := models.GetAllChatLatestMessege(username)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.AllChatResponse{
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
			c.JSON(http.StatusBadRequest, responses.PastChatResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, responses.PastChatResponse{
			Status:       http.StatusOK,
			Message:      `success!`,
			ChatMesseges: result,
		})
	}
}
