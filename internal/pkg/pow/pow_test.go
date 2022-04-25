package pow

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/mi7ter/go-word-of-wisdom/internal/pkg/pow/hasher"
)

func TestHashcashPow_Compute(t *testing.T) {
	type fields struct {
		maxIterations  int
		zeroHash       []rune
		hasher         hasher.Hasher
		expireDuration time.Duration
	}
	type args struct {
		resource string
		bits     int
		bytes    []byte
		date     time.Time
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Hashcash
		wantErr bool
	}{
		{
			"valid sha256 hashcash computation",
			fields{
				maxIterations:  1000000,
				zeroHash:       []rune("0000000000000000000000000000000000000000000000000000000000000000"),
				hasher:         hasher.NewSHA256(),
				expireDuration: time.Hour * 24 * 365,
			},
			args{
				resource: "test",
				bits:     5,
				bytes:    []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
				date:     time.Unix(1650312795, 0),
			},
			&Hashcash{
				Version:  1,
				Bits:     5,
				Date:     time.Unix(1650312795, 0),
				Resource: "test",
				Rand:     []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
				Counter:  43197,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Pow{
				maxIterations:  tt.fields.maxIterations,
				zeroHash:       tt.fields.zeroHash,
				hasher:         tt.fields.hasher,
				expireDuration: tt.fields.expireDuration,
			}
			got, err := h.Compute(tt.args.resource, tt.args.bits, tt.args.bytes, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Compute() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashcashPow_Verify(t *testing.T) {
	type fields struct {
		maxIterations  int
		zeroHash       []rune
		hasher         hasher.Hasher
		expireDuration time.Duration
	}
	type args struct {
		hash     string
		resource string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"valid sha256 hashcash verification",
			fields{
				maxIterations:  1000000,
				zeroHash:       []rune("0000000000000000000000000000000000000000000000000000000000000000"),
				hasher:         hasher.NewSHA256(),
				expireDuration: time.Hour * 24 * 365,
			},
			args{
				hash:     "1:5:1650312795:test::AQIDBAUGBwgJCgsMDQ4PEA==:YThiZA==",
				resource: "test",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Pow{
				maxIterations:  tt.fields.maxIterations,
				zeroHash:       tt.fields.zeroHash,
				hasher:         tt.fields.hasher,
				expireDuration: tt.fields.expireDuration,
			}
			if err := h.Verify(tt.args.hash, tt.args.resource); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashcashPow_isCorrectHash(t *testing.T) {
	type fields struct {
		maxIterations  int
		zeroHash       []rune
		hasher         hasher.Hasher
		expireDuration time.Duration
	}
	type args struct {
		hash      string
		zeroCount int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"valid sha256 with 5 leading zeros",
			fields{
				maxIterations:  10000000,
				zeroHash:       []rune("0000000000000000000000000000000000000000000000000000000000000000"),
				hasher:         hasher.NewSHA256(),
				expireDuration: time.Hour * 24 * 365,
			},
			args{
				hash:      "00000d4d6cede815f192b5a5ee56a7b6eaecbdf1a11c28f95476558f66767be4",
				zeroCount: 5,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Pow{
				maxIterations:  tt.fields.maxIterations,
				zeroHash:       tt.fields.zeroHash,
				hasher:         tt.fields.hasher,
				expireDuration: tt.fields.expireDuration,
			}
			if got := h.isCorrectHash(tt.args.hash, tt.args.zeroCount); got != tt.want {
				t.Errorf("isCorrectHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHashcashPow(t *testing.T) {
	type args struct {
		maxIterations  int
		h              hasher.Hasher
		expireDuration time.Duration
	}
	h := hasher.NewSHA256()
	tests := []struct {
		name string
		args args
		want *Pow
	}{
		{
			"hashcash instance",
			args{
				maxIterations:  1000000,
				h:              h,
				expireDuration: time.Hour * 24 * 365,
			},
			&Pow{
				maxIterations:  1000000,
				zeroHash:       []rune(strings.Repeat(string(zero), maxHashLength)),
				hasher:         h,
				expireDuration: time.Hour * 24 * 365,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPow(tt.args.maxIterations, tt.args.h, tt.args.expireDuration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
