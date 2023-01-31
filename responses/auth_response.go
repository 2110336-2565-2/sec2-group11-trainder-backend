package responses

type RegisterResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}

type LoginResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
}
