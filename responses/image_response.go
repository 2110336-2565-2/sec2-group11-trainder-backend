package responses

type ImageResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}
