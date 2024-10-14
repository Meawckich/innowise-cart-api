package model

import (
	"time"
)

type BaseError struct {
	TimeStamp  time.Time `json:"timestamp"`
	Error      string    `json:"error"`
	StatusCode int       `json:"status"`
	Message    string    `json:"message"`
}

func NewBaseError(status int, err, msg string) *BaseError {
	return &BaseError{
		TimeStamp:  time.Now(),
		Error:      err,
		StatusCode: status, //base
		Message:    msg,
	}
}
