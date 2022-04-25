package client

import "net"

// Service is the interface for the Wisdom client.
type Service interface {
	ComputeHashcash(hashcash []byte) ([]byte, error)
	HandleConnection(conn net.Conn) (string, error)
}
