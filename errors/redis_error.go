package ce

type RedisError struct {
	*BaseError
}

func ErrorRedis(err error) error {
	return &RedisError{BaseError: &BaseError{err}}
}
