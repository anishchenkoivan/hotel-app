package apperrors

import "fmt"

type NotFoundError struct {
	message string
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{message: message}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("not found: %s", e.message)
}

type ParsingError struct {
	message string
}

func NewParsingError(message string) *ParsingError {
	return &ParsingError{message: message}
}

func (e *ParsingError) Error() string {
	return fmt.Sprintf("parsing: %s", e.message)
}

var (
	NotFoundErrorInstance = NewNotFoundError("instance not found")
	ParsingErrorInstance  = NewParsingError("instance parsing")
)
