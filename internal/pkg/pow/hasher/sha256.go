package hasher

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256 returns a new SHA256 hasher.
type SHA256 struct{}

// NewSHA256 returns a new SHA256 hasher.
func NewSHA256() *SHA256 {
	return &SHA256{}
}

// Hash returns the sha256 hash of the data
func (h *SHA256) Hash(data string) (string, error) {
	s := sha256.New()

	if _, err := s.Write([]byte(data)); err != nil {
		return "", err
	}

	return hex.EncodeToString(s.Sum(nil)), nil
}
