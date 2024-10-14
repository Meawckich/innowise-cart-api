package model

import (
	"fmt"
	"net/http"
)

type InvalidRequestPathError struct {
	*BaseError
	Path string `json:"path"`
}

type InvalidRequestBodyError struct {
	*BaseError
	InvalidBody string `json:"invalidBody"`
}

func (e InvalidRequestPathError) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.Path)
}

func (e InvalidRequestBodyError) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.InvalidBody)
}

func NewInvalidBodyError(msg, body string) *InvalidRequestBodyError {
	return &InvalidRequestBodyError{
		BaseError: NewBaseError(
			http.StatusBadRequest,
			"Invalid body",
			msg,
		),
		InvalidBody: body,
	}
}

func NewInvalidRequestPathError(msg string) *InvalidRequestPathError {
	return &InvalidRequestPathError{
		BaseError: NewBaseError(
			http.StatusBadRequest,
			"Invalid request",
			msg,
		),
	}
}
