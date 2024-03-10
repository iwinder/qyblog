package errors

import "fmt"

type withCode struct {
	err   error
	msg   string
	code  int
	cause error
	*stack
}

func New(message string) error {
	return &withCode{
		msg:   message,
		err:   fmt.Errorf(message),
		code:  0,
		stack: callers(),
	}
}
func NewError(err error) error {
	return &withCode{
		err:   err,
		code:  0,
		stack: callers(),
	}
}
func WithCode(code int, format string, args ...interface{}) error {
	return &withCode{
		err:   fmt.Errorf(format, args...),
		code:  code,
		stack: callers(),
	}
}
func WrapC(err error, code int, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return &withCode{
		err:   fmt.Errorf(format, args...),
		msg:   fmt.Sprintf(format, args...),
		code:  code,
		cause: err,
		stack: callers(),
	}
}
func (w withCode) Error() string {
	return fmt.Sprintf("%v", w)
}

func (w *withCode) Cause() error  { return w.cause }
func (w *withCode) Unwrap() error { return w.cause }

func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}

		if cause.Cause() == nil {
			break
		}

		err = cause.Cause()
	}
	return err
}
