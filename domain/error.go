package domain

import "errors"

var (
	ErrInternalServerError = errors.New("internal Server Error")

	ErrNotFound = errors.New("not found")
)
