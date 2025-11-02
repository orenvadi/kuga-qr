package api

import "errors"

var (
	invalidJwt            = "missing or invalid token"
	serverError           = "internal server error"
	ErrInvalidCredentials = errors.New("invalid credentials")
)
