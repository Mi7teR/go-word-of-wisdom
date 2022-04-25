package pow

import (
	"reflect"
	"testing"
	"time"
)

func TestHashcashFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *Hashcash
		wantErr bool
	}{
		{
			name: "valid hashcash",
			args: args{
				s: "1:5:1650312795:test::AQIDBAUGBwgJCgsMDQ4PEA==:MA==",
			},
			want: &Hashcash{
				Version:  1,
				Bits:     5,
				Date:     time.Unix(1650312795, 0),
				Resource: "test",
				Rand:     []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
				Counter:  0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashcashFromString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashcashFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashcashFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashcash_ExpireTime(t *testing.T) {
	type fields struct {
		Version  int
		Bits     int
		Date     time.Time
		Resource string
		Rand     []byte
		Counter  int64
	}
	type args struct {
		duration time.Duration
	}
	date := time.Now()
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Time
	}{
		{
			"valid time",
			fields{
				Version:  1,
				Bits:     5,
				Date:     date,
				Resource: "",
				Rand:     []byte{},
				Counter:  0,
			},
			args{
				duration: time.Second * 10,
			},
			date.Add(time.Second * 10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hashcash{
				Version:  tt.fields.Version,
				Bits:     tt.fields.Bits,
				Date:     tt.fields.Date,
				Resource: tt.fields.Resource,
				Rand:     tt.fields.Rand,
				Counter:  tt.fields.Counter,
			}
			if got := h.ExpireTime(tt.args.duration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpireTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashcash_String(t *testing.T) {
	type fields struct {
		Version  int
		Bits     int
		Date     time.Time
		Resource string
		Rand     []byte
		Counter  int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"valid hashcash string",
			fields{
				Version:  1,
				Bits:     5,
				Date:     time.Unix(1650312795, 0),
				Resource: "test",
				Rand:     []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
				Counter:  0,
			},
			"1:5:1650312795:test::AQIDBAUGBwgJCgsMDQ4PEA==:MA==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hashcash{
				Version:  tt.fields.Version,
				Bits:     tt.fields.Bits,
				Date:     tt.fields.Date,
				Resource: tt.fields.Resource,
				Rand:     tt.fields.Rand,
				Counter:  tt.fields.Counter,
			}
			if got := h.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashcash_parse(t *testing.T) {
	type fields struct {
		Version  int
		Bits     int
		Date     time.Time
		Resource string
		Rand     []byte
		Counter  int64
	}
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"parse hashcash string",
			fields{
				Version:  1,
				Bits:     5,
				Date:     time.Unix(1650312795, 0),
				Resource: "test",
				Rand:     []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
				Counter:  0,
			},
			args{
				s: "1:5:1650312795:test::AQIDBAUGBwgJCgsMDQ4PEA==:MA==",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hashcash{
				Version:  tt.fields.Version,
				Bits:     tt.fields.Bits,
				Date:     tt.fields.Date,
				Resource: tt.fields.Resource,
				Rand:     tt.fields.Rand,
				Counter:  tt.fields.Counter,
			}
			if err := h.parse(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewHashcash(t *testing.T) {
	type args struct {
		bits     int
		resource string
		bytes    []byte
		date     time.Time
	}
	tests := []struct {
		name string
		args args
		want *Hashcash
	}{
		{
			"hashcash",
			args{
				bits:     5,
				resource: "test",
				bytes:    []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
				date:     time.Unix(1650312795, 0),
			},
			&Hashcash{
				Version:  1,
				Bits:     5,
				Date:     time.Unix(1650312795, 0),
				Resource: "test",
				Rand:     []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
				Counter:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashcash(tt.args.bits, tt.args.resource, tt.args.bytes, tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashcash() = %v, want %v", got, tt.want)
			}
		})
	}
}
