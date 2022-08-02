package helpers

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Response represents a response from the API
func ResponseAPI(status int, message string, data interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
