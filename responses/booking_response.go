package responses

import "trainder-api/models"

type CreateBookingResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}

type GetBookingsResponse struct {
	Status   int              `json:"status"`
	Message  string           `json:"message,omitempty"`
	Bookings []models.Booking `json:"bookings,omitempty"`
}
