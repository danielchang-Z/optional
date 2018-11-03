package optional

// Optional represents an immutable object that may contain a non-nil reference to another object.
// Each instance of this type either contains a non-nil reference, or contains nothing (in which
// case we say thatthe reference is "absent"); it is never said to "contain nil".
type Optional interface {
	// IsPresent returns true if optional contains a non-nil value.
	IsPresent() bool
	// IsPresent returns true if optional contains nothing.
	IsAbsent() bool
	// Get returns the non-nil value if optional is present, otherwise panics an illegal status error.
	Get() interface{}
	// Or returns the non-nil value if optional is present, otherwise returns the defaultValue.
	Or(defaultValue interface{}) interface{}
	// OrNil returns the non-nil value if optional is present, otherwise returns the defaultValue.
	OrNil() interface{}
	// OrSupply returns the non-nil value if optional is present, otherwise returns the result supplied by supplier.
	OrSupply(supplier func() interface{}) interface{}
	// Substitute returns the substitutive optional if current optional is absent, otherwise returns the original.
	Substitute(subsitutation Optional) Optional
	// Transform returns the optional transformed by given function if current optional if present, otherwise returns absent.
	Transform(function func(interface{}) interface{}) Optional
	// IfPresent consumes the non-nil value through the consumer if optional is present.
	IfPresent(consumer func(interface{}))
	// IfPresentOrElse consumes the non-nil value through the consumer if optional is present, otherwise call the worker.
	IfPresentOrElse(consumer func(interface{}), worker func())
}

// Absent returns an optional instance which contains nothing.
func Absent() Optional {
	return _absent
}

// Of returns an optional instance containing the given non-nil value.
func Of(value interface{}) Optional {
	return &present{checkNotNil(value)}
}

// OfNilable returns an optional instance containing the given value when it is not nil, otherwise returns the absent.
func OfNilable(value interface{}) Optional {
	if value == nil {
		return _absent
	}
	return &present{value}
}
