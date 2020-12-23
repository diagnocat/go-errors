package errors

type withCode struct {
	cause error
	code  int
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied code.
// If err is nil, Wrap returns nil.
func WrapC(err error, code int) error {
	return &withCode{
		cause: err,
		code:  code,
	}
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message and code.
// If err is nil, Wrap returns nil.
func WrapMC(err error, message string, code int) error {
	return &withCode{
		cause: Wrap(err, message),
		code:  code,
	}
}

func (e *withCode) Error() string {
	return e.cause.Error()
}

func (e *withCode) Cause() error { return e.cause }

// Unwrap provides compatibility for Go 1.13 error chains.
func (e *withCode) Unwrap() error { return e.cause }

func (e *withCode) Code() int {
	return e.code
}

func Code(err error) int {
	type coder interface {
		Code() int
	}
	type causer interface {
		Cause() error
	}

	var (
		coded coder
		cause causer
		ok    bool
	)

	for err != nil {
		coded, ok = err.(coder)
		if ok {
			return coded.Code()
		}

		cause, ok = err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}

	return 0
}
