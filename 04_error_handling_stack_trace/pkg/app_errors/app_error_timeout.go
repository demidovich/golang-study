package app_errors

import (
	"errors"
	"time"
)

type AppErrTimeout struct {
	Time  time.Duration
	stack []byte
	Err   error
}

func (e *AppErrTimeout) Error() string {
	return e.Time.String() + ": " + e.Err.Error() + string(e.stack)
}

func NewErrorTimeout(message string, time time.Duration) *AppErrTimeout {
	return &AppErrTimeout{
		Time:  time,
		stack: newStackTrace(),
		Err:   errors.New(message),
	}
}
