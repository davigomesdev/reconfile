package errors

import (
	"net/http"
	"strings"
)

type ErrorType string

const (
	INTERNAL_ERROR             ErrorType = "INTERNAL_ERROR"
	BAD_REQUEST_ERROR          ErrorType = "BAD_REQUEST_ERROR"
	NOT_FOUND_ERROR            ErrorType = "NOT_FOUND_ERROR"
	CONFLICT_ERROR             ErrorType = "CONFLICT_ERROR"
	UNAUTHORIZED_EXCEPTION     ErrorType = "UNAUTHORIZED_EXCEPTION"
	INVALID_CREDENTIALS        ErrorType = "INVALID_CREDENTIALS"
	UNPROCESSABLE_ENTITY_ERROR ErrorType = "UNPROCESSABLE_ENTITY_ERROR"
)

type AppError struct {
	Code     int       `json:"code"`
	Type     ErrorType `json:"type"`
	Messages []string  `json:"messages"`
}

func (e *AppError) Error() string {
	return strings.Join(e.Messages, "; ")
}

func NewInternalError() *AppError {
	return &AppError{
		Code:     http.StatusInternalServerError,
		Type:     INTERNAL_ERROR,
		Messages: []string{"Internal server error."},
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code:     http.StatusBadRequest,
		Type:     BAD_REQUEST_ERROR,
		Messages: []string{message},
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:     http.StatusNotFound,
		Type:     NOT_FOUND_ERROR,
		Messages: []string{message},
	}
}

func NewConflictError(message string) *AppError {
	return &AppError{
		Code:     http.StatusConflict,
		Type:     CONFLICT_ERROR,
		Messages: []string{message},
	}
}

func NewInvalidCredentialsError() *AppError {
	return &AppError{
		Code:     http.StatusForbidden,
		Type:     INVALID_CREDENTIALS,
		Messages: []string{"Credenciais inválidas."},
	}
}

func NewUnauthorizedExceptionError() *AppError {
	return &AppError{
		Code:     http.StatusUnauthorized,
		Type:     UNAUTHORIZED_EXCEPTION,
		Messages: []string{"Token inválido ou não fornecido."},
	}
}

func NewUnprocessableEntityError(message []string) *AppError {
	return &AppError{
		Code:     http.StatusUnprocessableEntity,
		Type:     UNPROCESSABLE_ENTITY_ERROR,
		Messages: message,
	}
}
