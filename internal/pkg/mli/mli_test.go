package mli

import (
	"reflect"
	"testing"
)

func TestEncodeWithMLI(t *testing.T) {
	type args struct {
		message []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"TestEncodeWithMLI",
			args{[]byte("Hello World")},
			[]byte{0, 13, 72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeWithMLI(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeWithMLI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMLI(t *testing.T) {
	type args struct {
		mli *[]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"TestGetMLI",
			args{&[]byte{0, 13}},
			13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMLI(tt.args.mli); got != tt.want {
				t.Errorf("GetMLI() = %v, want %v", got, tt.want)
			}
		})
	}
}
