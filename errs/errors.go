package errs

import (
	"net/http"
)

type ApplicationError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NotFoundException(message string) *ApplicationError {
	return &ApplicationError{Code: http.StatusNotFound, Message: message}
}
func BadRequestException(message string) *ApplicationError {
	return &ApplicationError{Code: http.StatusBadRequest, Message: message}
}

func InternalError(message string) *ApplicationError {
	return &ApplicationError{Code: http.StatusInternalServerError, Message: message}
}

func ValidationException(message string) *ApplicationError {
	return &ApplicationError{Code: http.StatusUnprocessableEntity, Message: message}
}
