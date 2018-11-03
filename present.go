package optional

// Implementation of optional which contains a non-nil value.
type present struct {
	value interface{}
}

func (o *present) IsPresent() bool {
	return true
}

func (o *present) IsAbsent() bool {
	return false
}

func (o *present) Get() interface{} {
	return o.value
}

func (o *present) Or(defaultValue interface{}) interface{} {
	checkNotNil(defaultValue)
	return o.value
}

func (o *present) OrNil() interface{} {
	return o.value
}

func (o *present) OrSupply(supplier func() interface{}) interface{} {
	checkNotNil(supplier())
	return o.value
}

func (o *present) Substitute(subsitutation Optional) Optional {
	checkNotNil(subsitutation)
	return o
}

func (o *present) Transform(function func(interface{}) interface{}) Optional {
	checkNotNil(function)
	return &present{checkNotNilWithMessage(function(o.value), "the function passed to Optional.transform() must not return nil")}
}

func (o *present) IfPresent(consumer func(interface{})) {
	checkNotNil(consumer)
	consumer(o.value)
}

func (o *present) IfPresentOrElse(consumer func(interface{}), worker func()) {
	checkNotNil(consumer)
	checkNotNil(worker)
	consumer(o.value)
}
