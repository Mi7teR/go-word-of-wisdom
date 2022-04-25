package client

import (
	"fmt"
	"log"
	"net"

	"github.com/mi7ter/go-word-of-wisdom/internal/domain/entity"
	"github.com/mi7ter/go-word-of-wisdom/internal/pkg/mli"
	"github.com/mi7ter/go-word-of-wisdom/internal/pkg/pow"
)

// Service is the client service
type Service struct {
	pow HashcashPow
}

// NewService creates a new client service
func NewService(pow HashcashPow) *Service {
	return &Service{
		pow: pow,
	}
}

// ComputeHashcash computes the valid hashcash for the given hashcash
func (s *Service) ComputeHashcash(hashcash []byte) ([]byte, error) {
	hc, err := pow.HashcashFromString(string(hashcash))
	if err != nil {
		return nil, fmt.Errorf("invalid hashcash: %v", err)
	}

	hashcash, err = s.pow.Compute(hc.Resource, hc.Bits, hc.Rand, hc.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to compute hashcash: %v", err)
	}

	return hashcash, nil
}

// HandleConnection handles the connection
// returns wisdom from server
func (s *Service) HandleConnection(conn net.Conn) (string, error) {
	log.Println("Sending request challenge message to server")
	// Send message with request challenge to server
	message, err := entity.NewMessage(entity.RequestChallengeMessageType, nil).ToBytes()
	if err != nil {
		return "", fmt.Errorf("failed to marshal message: %v", err)
	}

	_, err = conn.Write(mli.EncodeWithMLI(message))
	if err != nil {
		return "", fmt.Errorf("failed to send request challenge: %v", err)
	}

	// Read response from server
	b := make([]byte, mli.Size)

	_, err = conn.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to read response length: %v", err)
	}

	length := mli.GetMLI(&b)
	message = make([]byte, length)

	log.Println("Reading response from server")

	_, err = conn.Read(message)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	mess, err := entity.FromBytes(message)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal message: %v", err)
	}

	if mess.Type != entity.RequestChallengeMessageType {
		return "", fmt.Errorf("invalid message type: %v", mess.Type)
	}

	log.Println("Received response challenge message from server")

	hashcash, err := s.ComputeHashcash(mess.Payload)
	if err != nil {
		return "", fmt.Errorf("failed to compute hashcash: %v", err)
	}

	// Send message with response challenge to server
	message, err = entity.NewMessage(entity.ResponseChallengeMessageType, hashcash).ToBytes()
	if err != nil {
		return "", fmt.Errorf("failed to marshal message: %v", err)
	}

	log.Println("Sending response challenge message to server")

	_, err = conn.Write(mli.EncodeWithMLI(message))
	if err != nil {
		return "", fmt.Errorf("failed to send response challenge: %v", err)
	}

	log.Println("Reading wisdom from server")

	// Read response from server
	b = make([]byte, mli.Size)

	_, err = conn.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to read response length: %v", err)
	}

	length = mli.GetMLI(&b)
	message = make([]byte, length)

	_, err = conn.Read(message)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	mess, err = entity.FromBytes(message)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal message: %v", err)
	}

	if mess.Type != entity.WisdomMessageType {
		return "", fmt.Errorf("invalid message type: %v", mess.Type)
	}

	return string(mess.Payload), nil
}
