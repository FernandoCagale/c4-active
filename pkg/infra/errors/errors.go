package errors

import (
	"net/http"

	"github.com/FernandoCagale/c4-active/pkg/entity"
)

//ResponseError struct
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//AddInternalServerError InternalServerError
func AddInternalServerError(message string) ResponseError {
	return ResponseError{http.StatusInternalServerError, message}
}

//AddNotFoundError NotFound
func AddNotFoundError(message string) ResponseError {
	return ResponseError{http.StatusNotFound, message}
}

//AddUnauthorizedError Unauthorized
func AddUnauthorizedError(message string) ResponseError {
	return ResponseError{http.StatusUnauthorized, message}
}

//AddBadRequestError BadRequest
func AddBadRequestError(message string) ResponseError {
	return ResponseError{http.StatusBadRequest, message}
}

//AddMethodNotAllowedError MethodNotAllowed
func AddMethodNotAllowedError(message string) ResponseError {
	return ResponseError{http.StatusMethodNotAllowed, message}
}

//Err format err
func Err(err error) ResponseError {
	message := err.Error()
	switch err {
	case entity.ErrNotFound:
		return AddNotFoundError(message)
	case entity.ErrUnauthorized:
		return AddUnauthorizedError(message)
	case entity.ErrInvalidPayload:
		return AddBadRequestError(message)
	default:
		return AddInternalServerError(message)
	}
}
