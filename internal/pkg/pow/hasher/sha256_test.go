package hasher

import (
	"reflect"
	"testing"
)

func TestNewSHA256(t *testing.T) {
	tests := []struct {
		name string
		want *SHA256
	}{
		{
			name: "new sha256",
			want: &SHA256{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSHA256(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSHA256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA256_Hash(t *testing.T) {
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
			name: "sha256 hash",
			args: args{
				data: "test",
			},
			want:    "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &SHA256{}
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
