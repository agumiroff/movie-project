package repository

import (
	"errors"
	"fmt"
)

// Enumeration of repository error codes.
var (
	ErrCodeUnknown        = errors.New("unknown error occurred")
	ErrCodeDuplicateEntry = errors.New("duplicate entry")
	ErrCodeNotFound       = errors.New("not found")
	ErrCodeInvalidInput   = errors.New("invalid input")
	ErrCodeDatabase       = errors.New("database error")
	ErrCodeTimeout        = errors.New("operation timed out")
)

type RepError struct {
	Code    string
	Message string
}

func (e RepError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}
