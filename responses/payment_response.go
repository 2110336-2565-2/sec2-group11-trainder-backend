package responses

import "trainder-api/models"

type CreatePaymentResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}

type RequestPayoutResponse struct {
	Status    int    `json:"status"`
	Message   string `json:"message,omitempty"`
	BookingID string `json:"id,omitempty"`
}

type BookingListResponse struct {
	Status   int              `json:"status"`
	Message  string           `json:"message,omitempty"`
	Bookings []models.Booking `json:"bookings,omitempty"`
}
