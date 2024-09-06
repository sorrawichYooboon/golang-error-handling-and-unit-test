package ce

type InvalidFormatError struct {
	*BaseError
}

func ErrorInvalidFormat(err error) error {
	return &InvalidFormatError{BaseError: &BaseError{err}}
}
