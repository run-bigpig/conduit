package errcode

import "errors"

var (
	ErrInvalidClientId = errors.New("Invalid client id.")
	ErrClientIdEmpty   = errors.New("Client id is required.")
	ErrUserCodeEmpty   = errors.New("user_code is required")
	ErrInvalidParams   = errors.New("Invalid params.")
	ErrCertInvalid     = errors.New("Invalid cert.")
)
