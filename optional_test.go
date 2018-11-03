package optional

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAbsent(t *testing.T) {
	tests := []struct {
		name string
		want Optional
	}{
		{"absent", _absent},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Absent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Absent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOf(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want Optional
	}{
		{"non-nil value", args{1}, &present{1}},
		{"nil value", args{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := func() interface{} {
				defer func() {
					if r := recover(); r != nil {
						fmt.Printf("Recovered in test, reason:%s\n", r)
					}
				}()
				return Of(tt.args.value)
			}(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Of() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOfNilable(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want Optional
	}{
		{"non-nil value", args{0}, &present{0}},
		{"nil value", args{}, _absent},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OfNilable(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OfNilable() = %v, want %v", got, tt.want)
			}
		})
	}
}
