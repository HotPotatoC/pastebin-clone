package repository

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrMismatchPassword   = errors.New("mismatch password")
)
