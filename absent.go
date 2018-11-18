package optional

import (
	"github.com/pkg/errors"
)

// Implementation of optional which contains nothing.
type absent struct{}

// Instance of absent.
var _absent = &absent{}

func (o *absent) IsPresent() bool {
	return false
}

func (o *absent) IsAbsent() bool {
	return true
}

func (o *absent) Get() interface{} {
	panic(errors.Wrap(&illegalStatusError{}, "Optional.get() cannot be called on an absent value"))
}

func (o *absent) Or(defaultValue interface{}) interface{} {
	return checkNotNilWithMessage(defaultValue, "use Optional.OrNil instead of Optional.Or(nil)")
}

func (o *absent) OrNil() interface{} {
	return nil
}

func (o *absent) OrSupply(supplier func() interface{}) interface{} {
	return checkNotNilWithMessage(supplier(), "use Optional.OrNil instead of supplier returns nil")
}

func (o *absent) Substitute(subsitutation Optional) Optional {
	return checkNotNil(subsitutation).(Optional)
}

func (o *absent) Transform(function func(interface{}) interface{}) Optional {
	checkNotNil(function)
	return Absent()
}

func (o *absent) IfPresent(consumer func(interface{})) {
	checkNotNil(consumer)
}

func (o *absent) IfPresentOrElse(consumer func(interface{}), worker func()) {
	checkNotNil(consumer)
	checkNotNil(worker)
	worker()
}
