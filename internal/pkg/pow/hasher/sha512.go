package hasher

import (
	"crypto/sha512"
	"encoding/hex"
)

// SHA512 is a hasher that implements the SHA512 hash algorithm.
type SHA512 struct{}

// NewSHA512 returns a new SHA512 hasher.
func NewSHA512() *SHA512 {
	return &SHA512{}
}

// Hash returns the sha512 hash of the given data.
func (h *SHA512) Hash(data string) (string, error) {
	s := sha512.New()

	if _, err := s.Write([]byte(data)); err != nil {
		return "", err
	}

	return hex.EncodeToString(s.Sum(nil)), nil
}
