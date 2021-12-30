package result

import "net/http"

type BaseResult struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success() *BaseResult {
	return &BaseResult{
		Success: true,
		Code:    200,
		Message: "success",
		Data:    nil,
	}
}

func SuccessWithData(data interface{}) *BaseResult {
	return &BaseResult{
		Success: true,
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

func SuccessWithDataAndMessage(data interface{}, message string) *BaseResult {
	return &BaseResult{
		Success: true,
		Code:    200,
		Message: message,
		Data:    data,
	}
}

func Error() *BaseResult {
	return &BaseResult{
		Success: false,
		Code:    http.StatusBadRequest,
		Message: "error",
		Data:    nil,
	}
}

func ErrorWithMessage(message string) *BaseResult {
	return &BaseResult{
		Success: false,
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    nil,
	}
}
