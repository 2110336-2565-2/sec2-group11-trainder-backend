package controllers

import (
	"fmt"
	"net/http"
	"trainder-api/configs"
	"trainder-api/responses"

	"github.com/gin-gonic/gin"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type CreatePaymentForm struct {
	Token  string `json:"token" binding:"required"`
	Amount int64  `json:"amount" binding:"required"` // amount is temp should handle via booking id and calculate
}

//	@Summary		Create a payment
//	@Description	Create a payment using token and bookingId (currently taking amount)
//	@Tags			payment
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			body	CreatePayment	true	"details for creating payment"
//	@Success		200		{object}		responses.CreatePaymentResponse
//	@Router			/protected/create-payment [post]
func CreatePayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input CreatePaymentForm
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.CreateBookingResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		// Creates a charge from the token
		charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
			Amount:   input.Amount,
			Currency: "thb",
			Card:     input.Token,
		}

		if e := configs.OmiseClient.Do(charge, createCharge); e != nil {
			c.JSON(http.StatusBadRequest,
				responses.CreatePaymentResponse{
					Status:  http.StatusBadRequest,
					Message: e.Error(),
				})
			return
		}

		msg := fmt.Sprintf("charge: %s  amount: %s %d\n", charge.ID, charge.Currency, charge.Amount)
		c.JSON(http.StatusOK,
			responses.CreatePaymentResponse{
				Status:  http.StatusOK,
				Message: msg,
			})
	}
}
