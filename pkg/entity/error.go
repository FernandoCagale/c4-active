package entity

import (
	"errors"
)

//ErrNotFound not found
var ErrNotFound = errors.New("Not found")

// ErrUnauthorized unauthorized
var ErrUnauthorized = errors.New("Unauthorized")

//ErrInvalidPayload Invalid payload
var ErrInvalidPayload = errors.New("Invalid payload")

//ErrInternalServer Invalid payload
var ErrInternalServer = errors.New("Internal server error")
