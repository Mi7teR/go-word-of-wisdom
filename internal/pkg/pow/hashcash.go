package pow

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math"
	"math/big"
	mathRand "math/rand"
	"strconv"
	"strings"
	"time"
)

// Hashcash is a proof-of-work implementation that uses a hashcash-like
// algorithm.
type Hashcash struct {
	Version  int
	Bits     int
	Date     time.Time
	Resource string
	Rand     []byte
	Counter  int64
}

// NewHashcash returns a new Hashcash instance.
func NewHashcash(bits int, resource string, bytes []byte, date time.Time) *Hashcash {
	return &Hashcash{
		Version:  1,
		Bits:     bits,
		Date:     date,
		Resource: resource,
		Rand:     bytes,
		Counter:  0,
	}
}

// HashcashFromString returns a new Hashcash instance from a string.
func HashcashFromString(s string) (*Hashcash, error) {
	h := &Hashcash{}
	if err := h.parse(s); err != nil {
		return nil, err
	}

	return h, nil
}

// String returns the string representation of the hashcash.
func (h *Hashcash) String() string {
	return fmt.Sprintf("%d:%d:%d:%s::%s:%s",
		h.Version,
		h.Bits,
		h.Date.Unix(),
		h.Resource,
		base64.StdEncoding.EncodeToString(h.Rand),
		base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(h.Counter, 16))))
}

// Parse - parse the string representation of the hashcash.
func (h *Hashcash) parse(s string) error {
	parts := strings.Split(s, ":")
	if len(parts) != 7 {
		return fmt.Errorf("invalid hashcash string")
	}

	var err error

	h.Version, err = strconv.Atoi(parts[0])
	if err != nil {
		return err
	}

	h.Bits, err = strconv.Atoi(parts[1])
	if err != nil {
		return err
	}

	unixTime, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		return err
	}

	h.Date = time.Unix(unixTime, 0)
	h.Resource = parts[3]

	h.Rand, err = base64.StdEncoding.DecodeString(parts[5])
	if err != nil {
		return err
	}

	counterString, err := base64.StdEncoding.DecodeString(parts[6])
	if err != nil {
		return err
	}

	h.Counter, err = strconv.ParseInt(string(counterString), 16, 64)
	if err != nil {
		return err
	}

	return nil
}

// ExpireTime returns the time when the hashcash expires.
func (h Hashcash) ExpireTime(duration time.Duration) time.Time {
	return h.Date.Add(duration)
}

// RandomBytes returns a random byte slice
func RandomBytes() []byte {
	b, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		b = big.NewInt(mathRand.Int63n(math.MaxInt64)) //nolint:gosec
	}

	return b.Bytes()
}
