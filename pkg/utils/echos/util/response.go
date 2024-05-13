package util

type ResponseSucc struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponseFailed struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Cause   interface{} `json:"cause,omitempty"`
}

var WrapErrorResponse = func(message string, cause interface{}) ResponseFailed {
	return ResponseFailed{
		Status:  false,
		Message: message,
		Cause:   cause,
	}
}

var WrapSuccessResponse = func(message string, data interface{}) ResponseSucc {
	return ResponseSucc{
		Status:  true,
		Message: message,
		Data:    data,
	}
}
