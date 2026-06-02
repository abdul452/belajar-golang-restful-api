package exception

type NotFoundError struct {
	ErrorMessage string
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{
		ErrorMessage: message,
	}
}
