package ce

type InternalError struct {
	*BaseError
}

func ErrorInternal(err error) error {
	return &InternalError{BaseError: &BaseError{err}}
}
