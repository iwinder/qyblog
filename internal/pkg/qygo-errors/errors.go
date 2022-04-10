package qygo_errors

import "fmt"

type withCode struct {
	err   error
	code  int
	cause error
	*stack
}

func (w *withCode) Error() string {
	return fmt.Sprintf("%v", w)
}

func (w *withCode) Cause() error  { return w.cause }
func (w *withCode) Unwrap() error { return w.cause }

type withMessage struct {
	cause error
	msg   string
}

func (w *withMessage) Error() string {
	return w.msg
}

func (w *withMessage) Cause() error  { return w.cause }
func (w *withMessage) Unwrap() error { return w.cause }

type withStack struct {
	error
	*stack
}

func (w *withStack) Cause() error { return w.error }

func (w *withStack) Unwrap() error {
	if e, ok := w.error.(interface{ Unwrap() error }); ok {
		return e.Unwrap()
	}

	return w.error
}

func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(*withCode); ok {
		return &withCode{
			err:   fmt.Errorf(message),
			code:  e.code,
			cause: err,
			stack: callers(),
		}
	}
	err = &withMessage{
		cause: err,
		msg:   message,
	}

	return &withStack{err, callers()}
}

func WithCode(code int, format string, args ...interface{}) error {
	return &withCode{
		err:   fmt.Errorf(format, args...),
		code:  code,
		stack: callers(),
	}
}
