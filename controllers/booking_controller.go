package controllers

import (
	"net/http"
	"trainder-api/models"
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

// @Summary Create a new booking
// @Description Creates a new booking with the specified trainer, trainee, date, start time, and end time
// @Tags bookings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param json_in_ginContext body BookingForm true "put booking details and pass to gin.Context"
// @Success 200 {object} string "booking created successfully"
// @Failure 400 {object} string "bad request"
// @Failure 500 {object} string "internal server error"
// @Router /protected/create-booking [post]
func Book() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input BookingForm
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.CreateBookingResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.CreateBookingResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		err = models.CreateBooking(username, input.Trainer, input.Date, input.StartTime, input.EndTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.CreateBookingResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK,
			responses.CreateBookingResponse{
				Status:  http.StatusOK,
				Message: `success!`,
			})
	}
}
