package controllers

import (
	"fmt"
	"net/http"
	"trainder-api/configs"
	"trainder-api/models"
	"trainder-api/responses"
	"trainder-api/utils/tokens"

	"github.com/gin-gonic/gin"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type CreatePaymentForm struct {
	Token     string `json:"token" binding:"required"`
	BookingID string `json:"bookingID" binding:"required"` // amount is temp should handle via booking id and calculate
}

// @Summary		Create a payment
// @Description	Create a payment using token and bookingId
// @Tags			payment
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param 		input 	body 			CreatePaymentForm	true	"details for creating payment"
// @Success		200		{object}		responses.CreatePaymentResponse
// @Router			/protected/create-payment [post]
func CreatePayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input CreatePaymentForm
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.CreatePaymentResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, responses.CreatePaymentResponse{
				Status:  http.StatusUnauthorized,
				Message: `Cannot extract username from token`,
			})
			return
		}

		paymentInfo, err := models.GetPaymentInfo(input.BookingID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CreatePaymentResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		// Check if the user is really trainee
		if username != paymentInfo.TraineeUsername {
			c.JSON(http.StatusUnauthorized, responses.CreatePaymentResponse{
				Status:  http.StatusUnauthorized,
				Message: `booking can only be paid for by the trainee`,
			})
			return
		}
		if paymentInfo.BookingStatus != "confirm" {
			c.JSON(http.StatusBadRequest, responses.CreatePaymentResponse{
				Status:  http.StatusBadRequest,
				Message: `can only pay for booking that is confirmed`,
			})
			return
		}

		// Check if the booking is already paid
		if paymentInfo.PaymentStatus == "paid" {
			c.JSON(http.StatusBadRequest, responses.CreatePaymentResponse{
				Status:  http.StatusBadRequest,
				Message: `booking already paid`,
			})
			return
		}

		// Creates a charge from the token
		charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
			Amount:   paymentInfo.TotalCost * 100,
			Currency: "thb",
			Card:     input.Token,
		}

		if err := configs.OmiseClient.Do(charge, createCharge); err != nil {
			c.JSON(http.StatusBadRequest,
				responses.CreatePaymentResponse{
					Status:  http.StatusBadRequest,
					Message: err.Error(),
				})
			return
		}
		if charge.Status != omise.ChargeSuccessful {
			c.JSON(http.StatusOK,
				responses.CreatePaymentResponse{
					Status:  http.StatusOK,
					Message: fmt.Sprintf("charge fail with status: %s", charge.Status),
				})
			return

		}

		err = models.Pay(input.BookingID, charge.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.CreatePaymentResponse{
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
				})
			return

		}

		msg := fmt.Sprintf("booking: %s charge: %s  amount: %s %d.%02d\n", input.BookingID, charge.ID, charge.Currency, charge.Amount/100, charge.Amount%100)
		c.JSON(http.StatusOK,
			responses.CreatePaymentResponse{
				Status:  http.StatusOK,
				Message: msg,
			})
	}
}
