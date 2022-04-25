package server

import "time"

// WisdomRepository is the interface for the wisdom repository
type WisdomRepository interface {
	GetWisdom() string
}

// HashcashPow is the interface for the hashcash proof of work algorithm
type HashcashPow interface {
	GetNewHashcash(resource string, bits int, bytes []byte, date time.Time) ([]byte, error)
	Verify(hashcash, resource string) (bool, error)
}
