package apperror

import "net/http"

type Error struct {
	Code       string         `json:"code"`
	Message    string         `json:"message"`
	RequestId  string         `json:"requestId"`
	Details    map[string]any `json:"details"`
	StatusCode int            `json:"statusCode"`
}

func (e Error) Error() string {
	return e.Message
}

func (e *Error) ErrorCode(code string) {
	e.Code = code
}

const (
	InvalidFieldErrorCode = "invalid_field"
	ServerErrorCode       = "server_error"
)

func NewError(message string) Error {
	return Error{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Details:    map[string]any{},
	}
}

func NewInvalidFieldError(message string) Error {
	return Error{
		Code:       InvalidFieldErrorCode,
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Details:    map[string]any{},
	}
}

func NewServerError(message string) Error {
	return Error{
		Code:       ServerErrorCode,
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Details:    map[string]any{},
	}
}

func NewServiceUnavailableError(message string) Error {
	return Error{
		Message:    message,
		StatusCode: http.StatusServiceUnavailable,
		Details:    map[string]any{},
	}
}

func NewNotFoundError(message string) Error {
	return Error{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Details:    map[string]any{},
	}
}

func NewUnauthorizedError(message string) Error {
	return Error{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
		Details:    map[string]any{},
	}
}

func NewForbiddenError(message string) Error {
	return Error{
		Message:    message,
		StatusCode: http.StatusForbidden,
		Details:    map[string]any{},
	}
}

func NewTooManyRequestError(message string) Error {
	return Error{
		Message:    message,
		StatusCode: http.StatusTooManyRequests,
		Details:    map[string]any{},
	}
}

func NewGateWayTimeOutError(message string) Error {
	return Error{
		Message:    message,
		StatusCode: http.StatusGatewayTimeout,
		Details:    map[string]any{},
	}
}

func NewRequestTooLargeError(message string) Error {
	return Error{
		Message:    message,
		StatusCode: http.StatusRequestEntityTooLarge,
		Details:    map[string]any{},
	}
}

func NewServiceError(code, message string) Error {
	return Error{
		Code:       code,
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Details:    map[string]any{},
	}
}

func NewBadRequestError(message string) Error {
	return Error{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Details:    map[string]any{},
	}
}
