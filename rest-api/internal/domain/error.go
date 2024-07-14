package domain

// Represents an error occured during any operation.
type DomainError struct {
	Message string
	Code    int
}

func NewDomainError(code int, message string) error {
	return &DomainError{message, code}
}

func (e DomainError) Error() string {
	return e.Message
}
