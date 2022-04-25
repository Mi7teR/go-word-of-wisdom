package client

import (
	"time"
)

// HashcashPow is the interface for the hashcash pow for client
type HashcashPow interface {
	Compute(resource string, bits int, bytes []byte, date time.Time) ([]byte, error)
}
