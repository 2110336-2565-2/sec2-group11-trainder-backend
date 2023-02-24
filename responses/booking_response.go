package responses

type CreateBookingResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}

type UpdateBookingResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}

type DeleteBookingResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}
