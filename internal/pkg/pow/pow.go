package pow

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/mi7ter/go-word-of-wisdom/internal/pkg/pow/hasher"
)

const (
	zero          rune = 48 // 0 in ASCII
	maxHashLength      = 128
)

// error messages
var (
	ErrInvalidResource       = errors.New("invalid resource")
	ErrExpired               = errors.New("challenge expired")
	ErrInvalidHash           = errors.New("invalid hash")
	ErrMaxIterationsExceeded = errors.New("max iterations exceeded")
)

// Pow - simple implementation of hashcash PoW algorithm
type Pow struct {
	maxIterations  int
	zeroHash       []rune
	hasher         hasher.Hasher
	expireDuration time.Duration
}

// NewPow - creates new instance of Pow
func NewPow(maxIterations int, h hasher.Hasher, expireDuration time.Duration) *Pow {
	return &Pow{
		maxIterations:  maxIterations,
		zeroHash:       []rune(strings.Repeat(string(zero), maxHashLength)),
		hasher:         h,
		expireDuration: expireDuration,
	}
}

// Compute - compute hashcash PoW
func (h *Pow) Compute(resource string, bits int, bytes []byte, date time.Time) (*Hashcash, error) {
	if h.maxIterations > 0 {
		hashcash := NewHashcash(bits, resource, bytes, date)
		for int(hashcash.Counter) <= h.maxIterations {
			hashString, err := h.hasher.Hash(hashcash.String())
			if err != nil {
				return nil, err
			}

			if h.isCorrectHash(hashString, bits) {
				return hashcash, nil
			}

			hashcash.Counter++
		}
	}

	return nil, ErrMaxIterationsExceeded
}

// Verify - verify proof of work
func (h *Pow) Verify(hash string, resource string) error {
	hashcash, err := HashcashFromString(hash)
	if err != nil {
		return err
	}

	if hashcash.Resource != resource {
		return ErrInvalidResource
	}

	if time.Now().After(hashcash.ExpireTime(h.expireDuration)) {
		return ErrExpired
	}

	hashString, err := h.hasher.Hash(hashcash.String())
	if err != nil {
		return fmt.Errorf("failed to hash: %v", err)
	}

	if !h.isCorrectHash(hashString, hashcash.Bits) {
		return ErrInvalidHash
	}

	return nil
}

func (h *Pow) isCorrectHash(hash string, zeroCount int) bool {
	if zeroCount < len(hash) || zeroCount < len(h.zeroHash) {
		return strings.Compare(hash[:zeroCount], string(h.zeroHash[:zeroCount])) == 0
	}

	return false
}
