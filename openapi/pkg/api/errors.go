package api

import "errors"

var (
	invalidJwt            = "missing or invalid token"
	ErrInvalidCredentials = errors.New("invalid credentials")
)
