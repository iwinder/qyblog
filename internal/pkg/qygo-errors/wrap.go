package qygo_errors

import (
	goerrors "errors"
)

func Is(err, target error) bool {
	return goerrors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return goerrors.As(err, target)
}

func Unwrap(err error) error {
	return goerrors.Unwrap(err)
}
