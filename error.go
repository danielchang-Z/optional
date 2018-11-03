package optional

import (
	"github.com/pkg/errors"
)

var (
	// ErrIllegalStatus signals that a method has been invoked at an illegal or inappropriate time.
	ErrIllegalStatus = errors.New("illegal status")
	// ErrNilPointer indicates error is happened when an application attempts to use nil in a case where an object is required.
	ErrNilPointer = errors.New("nil pointer")
)

func checkNotNil(value interface{}) interface{} {
	if value == nil {
		panic(ErrNilPointer)
	}
	return value
}

func checkNotNilWithMessage(value interface{}, errorMessage string) interface{} {
	if value == nil {
		panic(errors.Wrap(ErrNilPointer, errorMessage))
	}
	return value
}
