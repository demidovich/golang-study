package app_errors

import "errors"

type AppError struct {
	stack []byte
	Err   error
}

func (e *AppError) Error() string {
	return errorWithStack(e.Err.Error(), e.stack)
}

func NewAppError(message string) *AppError {
	return &AppError{
		stack: newStackTrace(),
		Err:   errors.New(message),
	}
}
