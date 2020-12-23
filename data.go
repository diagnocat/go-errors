package errors

type withData struct {
	cause error
	data  interface{}
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied data.
// If err is nil, Wrap returns nil.
func WrapD(err error, data interface{}) error {
	return &withData{
		cause: err,
		data:  data,
	}
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message and data.
// If err is nil, Wrap returns nil.
func WrapMD(err error, message string, data interface{}) error {
	return &withData{
		cause: Wrap(err, message),
		data:  data,
	}
}

func (e *withData) Error() string {
	return e.cause.Error()
}

func (e *withData) Cause() error { return e.cause }

// Unwrap provides compatibility for Go 1.13 error chains.
func (e *withData) Unwrap() error { return e.cause }

func (e *withData) Data() interface{} {
	return e.data
}

func Data(err error) interface{} {
	type dataKeeper interface {
		Data() interface{}
	}
	type causer interface {
		Cause() error
	}

	var (
		keeper dataKeeper
		cause  causer
		ok     bool
	)

	for err != nil {
		keeper, ok = err.(dataKeeper)
		if ok {
			return keeper.Data()
		}

		cause, ok = err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}

	return nil
}
