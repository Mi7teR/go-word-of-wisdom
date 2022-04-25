package hashcashpow

import (
	"time"

	"github.com/mi7ter/go-word-of-wisdom/internal/pkg/pow"
	"github.com/mi7ter/go-word-of-wisdom/internal/pkg/pow/hasher"
)

// Repository - hashcash repository
type Repository struct {
	pow *pow.Pow
}

// Compute - compute hashcash
func (h *Repository) Compute(resource string, bits int, bytes []byte, date time.Time) ([]byte, error) {
	hashcash, err := h.pow.Compute(resource, bits, bytes, date)
	if err != nil {
		return nil, err
	}
	return []byte(hashcash.String()), nil
}

// NewHashcashPow return new hashcash proof of work repository
func NewHashcashPow(maxIterations int, hasher hasher.Hasher, expireDuration time.Duration) *Repository {
	return &Repository{
		pow: pow.NewPow(maxIterations, hasher, expireDuration),
	}
}

// GetNewHashcash - creates new hashcash
func (h *Repository) GetNewHashcash(resource string, bits int, bytes []byte, date time.Time) ([]byte, error) {
	hashcash := pow.NewHashcash(bits, resource, bytes, date)
	return []byte(hashcash.String()), nil
}

// Verify - verifies hashcash
func (h *Repository) Verify(hashcash, resource string) (bool, error) {
	err := h.pow.Verify(hashcash, resource)
	if err != nil {
		return false, err
	}

	return true, nil
}
