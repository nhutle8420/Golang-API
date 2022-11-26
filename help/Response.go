package help

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// EmptyObj được sử dụng khi không muốn dữ liệu trả về null trên json
type EmptyObj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	t := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return t
}
func BuildErrorResponse(message string, err string, data interface{}) Response {
	x := strings.Split(err, "\n")
	t := Response{
		Status: false,
		Error:  x,
		Data:   data,
	}
	return t
}
