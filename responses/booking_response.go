package responses

import "trainder-api/models"

type CreateBookingResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}

type GetBookingsResponse struct {
	Status   int                    `json:"status"`
	Message  string                 `json:"message,omitempty"`
	Bookings []models.ReturnBooking `json:"bookings,omitempty"`
}

type GetBookingResponse struct {
	Status   int            `json:"status"`
	Message  string         `json:"message,omitempty"`
	Bookings models.Booking `json:"booking,omitempty"`
}
type UpdateBookingResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}
type DeleteBookingResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}
