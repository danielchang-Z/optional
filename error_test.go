package optional

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func TestIsRuntimeError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"illegal status error", args{&illegalStatusError{}}, true},
		{"nil pointer error", args{&nilPointerError{}}, true},
		{"other error", args{fmt.Errorf("other")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRuntimeError(tt.args.err); got != tt.want {
				t.Errorf("IsRuntimeError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCausedByRuntimeError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"caused by illegal status error", args{errors.Wrap(&illegalStatusError{}, "test")}, true},
		{"caused by nil pointer error", args{errors.Wrap(&nilPointerError{}, "test")}, true},
		{"caused by other error", args{errors.Wrap(fmt.Errorf("other"), "test")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCausedByRuntimeError(tt.args.err); got != tt.want {
				t.Errorf("IsCausedByRuntimeError() = %v, want %v", got, tt.want)
			}
		})
	}
}
