package responses

type CreatePaymentResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}
