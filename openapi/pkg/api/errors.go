package api

import "errors"

var (
	scheduleNotFound      = "schedule not found"
	invalidJwt            = "missing or invalid token"
	serverError           = "internal server error"
	unauthorized          = "user is not authorized"
	ErrInvalidCredentials = errors.New("invalid credentials")
)
