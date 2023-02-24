package responses

type CreateBookingResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}