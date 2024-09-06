package ce

type InvalidRequestError struct {
	*BaseError
}

func ErrorInvalidRequest(err error) error {
	return &InvalidRequestError{BaseError: &BaseError{err}}
}
