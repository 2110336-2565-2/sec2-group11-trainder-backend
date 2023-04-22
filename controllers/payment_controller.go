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

type RequestPayoutForm struct {
	BookingID     string `json:"bookingID" binding:"required"`
	Bank          string `json:"bank" binding:"required"`
	AccountNumber string `json:"accountNumber" binding:"required"`
	AccountName   string `json:"accountName" binding:"required"`
}

// @Summary		Request a payout
// @Description	Mark payment as needed payout
// @Tags			payment
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param 		input 	body 			RequestPayoutForm	true	"details for requesting payout"
// @Success		200		{object}		responses.RequestPayoutResponse
// @Success		400		{object}		responses.RequestPayoutResponse
// @Success		401		{object}		responses.RequestPayoutResponse
// @Success		403		{object}		responses.RequestPayoutResponse
// @Router			/protected/request-payout [post]
func RequestPayout() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input RequestPayoutForm
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.RequestPayoutResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		username, err := tokens.ExtractTokenUsername(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, responses.RequestPayoutResponse{
				Status:  http.StatusUnauthorized,
				Message: `Cannot extract username from token`,
			})
			return
		}

		paymentInfo, err := models.GetPaymentInfo(input.BookingID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.RequestPayoutResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		if username != paymentInfo.TrainerUsername {
			c.JSON(http.StatusForbidden, responses.RequestPayoutResponse{
				Status:  http.StatusForbidden,
				Message: `only trainer can request payout`,
			})
			return
		}
		if paymentInfo.PaymentStatus != "paid" {
			c.JSON(http.StatusForbidden, responses.RequestPayoutResponse{
				Status:  http.StatusForbidden,
				Message: `only paid payment can be payout`,
			})
			return
		}
		err = models.RequestPayout(input.BookingID, input.Bank, input.AccountName, input.AccountNumber)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.RequestPayoutResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK,
			responses.RequestPayoutResponse{
				Status:    http.StatusOK,
				Message:   "success",
				BookingID: input.BookingID,
			})

	}
}

type PayoutForm struct {
	BookingID string `json:"bookingID" binding:"required"` // amount is temp should handle via booking id and calculate
}

// @Summary		Payout
// @Description	Mark payment as payout
// @Tags			payment
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param 		input 	body 			PayoutForm	true	"details for payout"
// @Success		200		{object}		responses.RequestPayoutResponse
// @Success		400		{object}		responses.RequestPayoutResponse
// @Success		401		{object}		responses.RequestPayoutResponse
// @Success		403		{object}		responses.RequestPayoutResponse
// @Router			/protected/payout [post]
func Payout() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input RequestPayoutForm
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.RequestPayoutResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		username, err := tokens.ExtractTokenUsername(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, responses.RequestPayoutResponse{
				Status:  http.StatusUnauthorized,
				Message: `Cannot extract username from token`,
			})
			return
		}

		if !models.IsAdmin(username) {
			c.JSON(http.StatusForbidden, responses.RequestPayoutResponse{
				Status:  http.StatusForbidden,
				Message: `only admin allowed`,
			})
			return
		}

		paymentInfo, err := models.GetPaymentInfo(input.BookingID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.RequestPayoutResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		if paymentInfo.PaymentStatus != "need_payout" {
			c.JSON(http.StatusBadRequest, responses.RequestPayoutResponse{
				Status:  http.StatusBadRequest,
				Message: `only need_payout payment can be paid out`,
			})
			return
		}
		err = models.Payout(input.BookingID)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.RequestPayoutResponse{
				Status:  http.StatusForbidden,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK,
			responses.RequestPayoutResponse{
				Status:    http.StatusOK,
				Message:   "success",
				BookingID: input.BookingID,
			})

	}

}

// @Summary		Get Payment list
// @Description	Get Payment list for trainer
// @Tags			payment
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Success		200		{object}		responses.BookingListResponse
// @Success		400		{object}		responses.BookingListResponse
// @Success		401		{object}		responses.BookingListResponse
// @Success		403		{object}		responses.BookingListResponse
// @Router			/protected/payment-list [get]
func PaymentList() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := tokens.ExtractTokenUsername(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, responses.BookingListResponse{
				Status:  http.StatusUnauthorized,
				Message: `Cannot extract username from token`,
			})
			return
		}
		if !models.IsTrainer(username) {
			c.JSON(http.StatusForbidden, responses.BookingListResponse{
				Status:  http.StatusForbidden,
				Message: `only trainer can request payout`,
			})
			return
		}
		payments, err := models.GetPaidBookings(username)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookingListResponse{
				Status:  http.StatusUnauthorized,
				Message: err.Error(),
			})
			return

		}
		fmt.Println(payments)

		c.JSON(http.StatusOK, responses.BookingListResponse{
			Status:   http.StatusOK,
			Message:  `success`,
			Bookings: payments,
		})

	}
}

// @Summary		Get Payment Need Payout
// @Description	Get Payment list that is needed payout
// @Tags			payment
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Success		200		{object}		responses.BookingListResponse
// @Success		400		{object}		responses.BookingListResponse
// @Success		401		{object}		responses.BookingListResponse
// @Success		403		{object}		responses.BookingListResponse
// @Router			/protected/payment-need-payouts [get]
func PaymentNeedPayouts() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := tokens.ExtractTokenUsername(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, responses.BookingListResponse{
				Status:  http.StatusUnauthorized,
				Message: `Cannot extract username from token`,
			})
			return
		}
		if !models.IsAdmin(username) {
			c.JSON(http.StatusForbidden, responses.BookingListResponse{
				Status:  http.StatusForbidden,
				Message: `only admin allowed`,
			})
			return
		}

		payments, err := models.BookingNeedPayouts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookingListResponse{
				Status:  http.StatusUnauthorized,
				Message: err.Error(),
			})
			return

		}

		c.JSON(http.StatusOK, responses.BookingListResponse{
			Status:   http.StatusOK,
			Message:  `success`,
			Bookings: payments,
		})

	}
}
