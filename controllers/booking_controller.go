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

type UpdateBookingForm struct {
	BookingId     string `json:"bookingId" binding:"required"`
	Status        string `json:"status"`
	PaymentStatus string `json:"paymentStatus"`
}

type DeleteBookingForm struct {
	BookingId string `json:"bookingId" binding:"required"`
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

// @Summary Get bookings for the logged in user
// @Description Retrieve a list of upcoming bookings for the user who is currently logged in
// @Tags bookings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.GetBookingsResponse
// @Failure 400 {object} responses.GetBookingsResponse
// @Router /protected/bookings [GET]
func GetBookings() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.GetBookingsResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
		}
		result, err := models.GetUpcomingBookings(username)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.GetBookingsResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, responses.GetBookingsResponse{
			Status:   http.StatusOK,
			Message:  `success!`,
			Bookings: result,
		})
	}
}

// @Summary Get booking by ID
// @Description Retrieve a single booking using id
// @Tags bookings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.GetBookingResponse
// @Failure 400 {object} responses.GetBookingResponse
// @Router /protected/booking [GET]
func GetBooking() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, responses.GetBookingResponse{
				Status:  http.StatusUnauthorized,
				Message: err.Error(),
			})
		}
		queryParams := c.Request.URL.Query()
		bookingIDs, prs := queryParams["id"]
		if !prs {
			c.JSON(http.StatusBadRequest, responses.GetBookingResponse{
				Status:  http.StatusBadRequest,
				Message: "id not found in query",
			})
		}

		// Only use the first value in query
		bookingID := bookingIDs[0]
		result, err := models.GetBooking(bookingID)
		if result.Trainee != username && result.Trainer != username {
			c.JSON(http.StatusUnauthorized, responses.GetBookingResponse{
				Status:  http.StatusUnauthorized,
				Message: "can only view own booking",
			})
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, responses.GetBookingResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, responses.GetBookingResponse{
			Status:   http.StatusOK,
			Message:  `success!`,
			Bookings: result,
		})
	}
}

// @Summary Update a booking
// @Description Update a booking of sepecified bookingId with the specified update input consist of status(pending/confirm/complete) and paymentStatus(pending/paid)
// @Tags bookings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param json_in_ginContext body UpdateBookingForm true "put updateBookingForm details and pass to gin.Context"
// @Success	200		{object}	responses.UpdateBookingResponse	"Successfully update booking"
// @Failure	400		{object}	responses.UpdateBookingResponse	"Bad Request, missing filed of objectId or cannot find bookingObjectId"
// @Router /protected/update-booking [post]
func UpdateBooking() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input UpdateBookingForm
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.UpdateBookingResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		err := models.UpdateBooking(input.BookingId, input.Status, input.PaymentStatus)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.UpdateBookingResponse{
				Status:  http.StatusBadRequest,
				Message: `update failed ` + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK,
			responses.UpdateBookingResponse{
				Status:  http.StatusOK,
				Message: `update booking success!`,
			})
	}
}

// @Summary Delete a booking
// @Description Delete a booking with the specified bookingId
// @Tags bookings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param json_in_ginContext body DeleteBookingForm true "put DeleteBookingForm details and pass to gin.Context"
// @Success	200		{object}	responses.DeleteBookingResponse	"Successfully delete booking"
// @Failure	400		{object}	responses.DeleteBookingResponse	"Bad Request, missing filed of objectId or cannot find bookingObjectId"
// @Router /protected/delete-booking [delete]
func DeleteBooking() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input DeleteBookingForm
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.DeleteBookingResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		err := models.DeleteBooking(input.BookingId)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.DeleteBookingResponse{
				Status:  http.StatusBadRequest,
				Message: `delete booking failed: ` + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK,
			responses.UpdateBookingResponse{
				Status:  http.StatusOK,
				Message: `delete booking success!`,
			})
	}
}

// @Summary Get today bookings for the logged in user
// @Description Retrieve a list of today bookings for the user who is currently logged in
// @Tags bookings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.GetBookingsResponse
// @Failure 400 {object} responses.GetBookingsResponse
// @Router /protected/today-event [GET]
func GetTodayEvents() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.GetBookingsResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
		}

		result, err := models.GetTodayBookings(username)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.GetBookingsResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, responses.GetBookingsResponse{
			Status:   http.StatusOK,
			Message:  `success!`,
			Bookings: result,
		})
	}
}
