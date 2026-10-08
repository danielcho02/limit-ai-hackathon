package v1

// common response structure
type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}