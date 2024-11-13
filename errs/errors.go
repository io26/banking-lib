package errs

import "net/http"

type AppError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func NewNotFoundError(message string) *AppError {
	return appError(http.StatusNotFound, message)
}

func NewUnexpectedError(message string) *AppError {
	return appError(http.StatusInternalServerError, message)
}

func NewBadRequest(message string) *AppError {
	return appError(http.StatusBadRequest, message)
}

func NewValidationError(message string) *AppError {
	return appError(http.StatusUnprocessableEntity, message)
}

func NewUnauthorizedError(message string) *AppError {
	return appError(http.StatusUnauthorized, message)
}

func appError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

var InvalidAccountAmount = NewValidationError("To open a new account you need to deposit at least 5000.00")
var InvalidAccountType = NewValidationError("To open a new account you need to deposit at least 5000.00")

var InvalidTransactionAmount = NewValidationError("Transaction amount can't be negative")
var InvalidTransactionType = NewValidationError("Transaction type should be withdrawal or deposit")
