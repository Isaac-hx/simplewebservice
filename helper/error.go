package helper

import "fmt"

type CustomError struct {
	Message    string
	StatusCode int
	Cause      error
}

func (e *CustomError) Unwrap() error {
	return e.Cause
}

// Implementasi interface error
func (e *CustomError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}
func NewCustomError(message string, statusCode int, cause error) error {
	return &CustomError{
		Message:    message,
		StatusCode: statusCode,
		Cause:      cause,
	}
}

var (
	ErrInvalidDateType   = NewCustomError("Invalid date type", 400, nil)
	ErrInvalidCoverUrl   = NewCustomError("Invalid cover url", 400, nil)
	ErrEmptyTitle        = NewCustomError("Title can't be empty", 400, nil)
	ErrInvalidQueryParam = NewCustomError("Invalid query param", 400, nil)
	ErrInvalidPathValue  = NewCustomError("Invalid path value", 400, nil)
)
