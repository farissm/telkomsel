package helper

type Response struct {
	Code	int		`json:"code"`
	Message	string	`json:"message"`
}

func APIResponse(code int, message string) Response {
	jsonResponse := Response{
		Code: code,
		Message: message,
	}

	return jsonResponse
}
