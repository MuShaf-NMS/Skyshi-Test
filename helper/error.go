package helper

import "fmt"

type CustomError struct {
	Code    int
	Status  string
	Message string
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("%d %s; errors: %s", ce.Code, ce.Status, ce.Message)
}

// Helper to create custom error
func NewError(statusCode int, status string, msg string) error {
	return &CustomError{
		Code:    statusCode,
		Status:  status,
		Message: msg,
	}
}
