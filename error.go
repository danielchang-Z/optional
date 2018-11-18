package optional

import (
	"github.com/pkg/errors"
)

// The RuntimeError interface identifies a run time error.
type RuntimeError interface {
	error

	// runtimeError is a no-op function but serves to distinguish types
	// that are run time errors from ordinary errors.
	runtimeError()
}

// IsRuntimeError indicates whether an error is RuntimeError type.
func IsRuntimeError(err error) bool {
	_, ok := err.(RuntimeError)
	return ok
}

type illegalStatusError struct{}

func (e *illegalStatusError) Error() string {
	return "illegal status"
}

func (e *illegalStatusError) runtimeError() {}

type nilPointerError struct{}

func (e *nilPointerError) Error() string {
	return "nil point"
}

func (e *nilPointerError) runtimeError() {}

func checkNotNil(value interface{}) interface{} {
	return checkNotNilWithMessage(value, "")
}

func checkNotNilWithMessage(value interface{}, errorMessage string) interface{} {
	if value == nil {
		panic(errors.Wrap(&nilPointerError{}, errorMessage))
	}
	return value
}

// IsCausedByRuntimeError indicates whether an error is caused by RuntimeError.
func IsCausedByRuntimeError(err error) bool {
	return IsRuntimeError(errors.Cause(err))
}
