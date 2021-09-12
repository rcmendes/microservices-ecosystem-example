package web

type ErrorMessage struct {
	StatusCode int    `json:"status_code"`
	Context    string `json:"context"`
	Message    string `json:"message"`
}

func LoginError(statusCode int, err error) ErrorMessage {
	return ErrorMessage{
		StatusCode: statusCode,
		Context:    "authentication:login",
		Message:    err.Error(),
	}
}
