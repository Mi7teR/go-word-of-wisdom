package server

import (
	"errors"
	"time"

	"github.com/mi7ter/go-word-of-wisdom/internal/domain/entity"
	"github.com/mi7ter/go-word-of-wisdom/internal/pkg/pow"
)

// Service is a server.Service interface implementation.
type Service struct {
	hashcashPow HashcashPow
	repo        WisdomRepository
	bits        int
}

var (
	// ErrCloseMessage is returned when the service is closing connection.
	ErrCloseMessage = errors.New("close connection")
	// ErrInvalidMessage is returned when the message is invalid.
	ErrInvalidMessage = errors.New("invalid message")
	// ErrInvalidHashcash is returned when the hashcash is invalid.
	ErrInvalidHashcash = errors.New("invalid hashcash")
)

// NewService return new Service instance
func NewService(pow HashcashPow, repo WisdomRepository, bits int) *Service {
	return &Service{
		hashcashPow: pow,
		repo:        repo,
		bits:        bits,
	}
}

// ProcessMessage is as message processor.
func (s *Service) ProcessMessage(message []byte, resource string) ([]byte, error) {
	mess, err := entity.FromBytes(message)
	if err != nil {
		return nil, err
	}

	switch mess.Type { //nolint:exhaustive
	case entity.CloseMessageType:
		return nil, ErrCloseMessage
	case entity.RequestChallengeMessageType:
		payload, err := s.hashcashPow.GetNewHashcash(resource, s.bits, pow.RandomBytes(), time.Now())
		if err != nil {
			return nil, err
		}

		return entity.NewMessage(entity.RequestChallengeMessageType, payload).ToBytes()
	case entity.ResponseChallengeMessageType:
		hashcash := string(mess.Payload)

		ok, err := s.hashcashPow.Verify(hashcash, resource)
		if err != nil {
			return nil, err
		}

		if !ok {
			return nil, ErrInvalidHashcash
		}

		wisdom := s.repo.GetWisdom()

		return entity.NewMessage(entity.WisdomMessageType, []byte(wisdom)).ToBytes()
	default:
		return nil, ErrInvalidMessage
	}
}
