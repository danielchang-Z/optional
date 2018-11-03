package optional

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_present_IsPresent(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"return true", fields{1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &present{
				value: tt.fields.value,
			}
			if got := o.IsPresent(); got != tt.want {
				t.Errorf("present.IsPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_present_IsAbsent(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"return false", fields{1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &present{
				value: tt.fields.value,
			}
			if got := o.IsAbsent(); got != tt.want {
				t.Errorf("present.IsAbsent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_present_Get(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{"get immutable object in optional", fields{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &present{
				value: tt.fields.value,
			}
			if got := o.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("present.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_present_Or(t *testing.T) {
	type fields struct {
		value interface{}
	}
	type args struct {
		defaultValue interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{"get immutable object in optional", fields{1}, args{2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &present{
				value: tt.fields.value,
			}
			if got := o.Or(tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("present.Or() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_present_OrNil(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{"get immutable object in optional", fields{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &present{
				value: tt.fields.value,
			}
			if got := o.OrNil(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("present.OrNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_present_OrSupply(t *testing.T) {
	type fields struct {
		value interface{}
	}
	type args struct {
		supplier func() interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{"does not use supplier", fields{1}, args{func() interface{} { return 2 }}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &present{
				value: tt.fields.value,
			}
			if got := o.OrSupply(tt.args.supplier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("present.OrSupply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_present_Substitute(t *testing.T) {
	type fields struct {
		value interface{}
	}
	type args struct {
		secondChoice Optional
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Optional
	}{
		{"does not need to substitute", fields{1}, args{Of(2)}, Of(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &present{
				value: tt.fields.value,
			}
			if got := o.Substitute(tt.args.secondChoice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("present.Substitute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_present_Transform(t *testing.T) {
	type fields struct {
		value interface{}
	}
	type args struct {
		function func(interface{}) interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Optional
	}{
		{"transformation", fields{1}, args{func(x interface{}) interface{} { return fmt.Sprintf("%v", x) }}, Of("1")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &present{
				value: tt.fields.value,
			}
			if got := o.Transform(tt.args.function); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("present.Transform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_present_IfPresent(t *testing.T) {
	var tmp string
	type fields struct {
		value interface{}
	}
	type args struct {
		consumer func(interface{})
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"use consumer when present", fields{1}, args{func(x interface{}) { tmp = fmt.Sprintf("%v", x) }}, "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &present{
				value: tt.fields.value,
			}
			o.IfPresent(tt.args.consumer)
			if !reflect.DeepEqual(tmp, tt.want) {
				t.Errorf("absent.Transform() = %v, want %v", tmp, tt.want)
			}
		})
	}
}

func Test_present_IfPresentOrElse(t *testing.T) {
	var tmp string
	type fields struct {
		value interface{}
	}
	type args struct {
		consumer func(interface{})
		worker   func()
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"use consumer and ignore worker  when absent", fields{1}, args{
			func(x interface{}) { tmp = fmt.Sprintf("%v", x) },
			func() { tmp = "worker" },
		}, "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &present{
				value: tt.fields.value,
			}
			o.IfPresentOrElse(tt.args.consumer, tt.args.worker)
			if !reflect.DeepEqual(tmp, tt.want) {
				t.Errorf("absent.Transform() = %v, want %v", tmp, tt.want)
			}
		})
	}
}
