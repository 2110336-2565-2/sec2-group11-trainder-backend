package controllers

import (
	"net/http"
	"trainder-api/responses"
	"trainder-api/utils/tokens"

	"github.com/gin-gonic/gin"
)

type BookingForm struct {
	Trainer   string `json:"trainer"`
	Date      string `json:"date"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

func CreateBooking() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input BookingForm
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.ProfileResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ProfileResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		_ = username

		c.JSON(http.StatusOK,
			responses.ProfileResponse{
				Status:  http.StatusOK,
				Message: `success!`,
			})
	}
}
