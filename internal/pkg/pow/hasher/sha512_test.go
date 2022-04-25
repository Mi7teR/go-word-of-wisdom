package hasher

import (
	"reflect"
	"testing"
)

func TestNewSHA512(t *testing.T) {
	tests := []struct {
		name string
		want *SHA512
	}{
		{
			name: "new sha512",
			want: &SHA512{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSHA512(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSHA512() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA512_Hash(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "sha512 hash",
			args: args{
				data: "test",
			},
			want:    "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &SHA512{}
			got, err := h.Hash(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hash() got = %v, want %v", got, tt.want)
			}
		})
	}
}
