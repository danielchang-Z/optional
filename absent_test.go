package optional

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_absent_IsPresent(t *testing.T) {
	tests := []struct {
		name string
		o    *absent
		want bool
	}{
		{"return false", _absent, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &absent{}
			if got := o.IsPresent(); got != tt.want {
				t.Errorf("absent.IsPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_absent_IsAbsent(t *testing.T) {
	tests := []struct {
		name string
		o    *absent
		want bool
	}{
		{"return true", _absent, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &absent{}
			if got := o.IsAbsent(); got != tt.want {
				t.Errorf("absent.IsAbsent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_absent_Get(t *testing.T) {
	tests := []struct {
		name string
		o    *absent
		want interface{}
	}{
		{"recover from panic", _absent, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := func() interface{} {
				defer func() {
					if r := recover(); r != nil {
						fmt.Printf("Recovered in test, reason:%s\n", r)
					}
				}()
				return tt.o.Get()
			}(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("absent.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_absent_Or(t *testing.T) {
	type args struct {
		defaultValue interface{}
	}
	tests := []struct {
		name string
		o    *absent
		args args
		want interface{}
	}{
		{"use default value", _absent, args{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Or(tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("absent.Or() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_absent_OrNil(t *testing.T) {
	tests := []struct {
		name string
		o    *absent
		want interface{}
	}{
		{"return nil", _absent, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.OrNil(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("absent.OrNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_absent_OrSupply(t *testing.T) {
	type args struct {
		supplier func() interface{}
	}
	tests := []struct {
		name string
		o    *absent
		args args
		want interface{}
	}{
		{"use supplier", _absent, args{func() interface{} { return 1 }}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.OrSupply(tt.args.supplier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("absent.OrSupply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_absent_Substitute(t *testing.T) {
	type args struct {
		secondChoice Optional
	}
	tests := []struct {
		name string
		o    *absent
		args args
		want Optional
	}{
		{"substitution", _absent, args{Of(1)}, Of(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Substitute(tt.args.secondChoice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("absent.Substitute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_absent_Transform(t *testing.T) {
	type args struct {
		function func(interface{}) interface{}
	}
	tests := []struct {
		name string
		o    *absent
		args args
		want Optional
	}{
		{"transformation", _absent, args{func(x interface{}) interface{} { return fmt.Sprintf("%v", x) }}, _absent},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Transform(tt.args.function); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("absent.Transform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_absent_IfPresent(t *testing.T) {
	var tmp string
	type args struct {
		consumer func(interface{})
	}
	tests := []struct {
		name string
		o    *absent
		args args
		want string
	}{
		{"ignore consumer when absent", _absent, args{func(x interface{}) { tmp = fmt.Sprintf("%v", x) }}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.o.IfPresent(tt.args.consumer)
			if !reflect.DeepEqual(tmp, tt.want) {
				t.Errorf("absent.Transform() = %v, want %v", tmp, tt.want)
			}
		})
	}
}

func Test_absent_IfPresentOrElse(t *testing.T) {
	var tmp string
	type args struct {
		consumer func(interface{})
		worker   func()
	}
	tests := []struct {
		name string
		o    *absent
		args args
		want string
	}{
		{"use worker and ignore consumer when absent", _absent, args{
			func(x interface{}) { tmp = fmt.Sprintf("%v", x) },
			func() { tmp = "worker" },
		}, "worker"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.o.IfPresentOrElse(tt.args.consumer, tt.args.worker)
			if !reflect.DeepEqual(tmp, tt.want) {
				t.Errorf("absent.Transform() = %v, want %v", tmp, tt.want)
			}
		})
	}
}
