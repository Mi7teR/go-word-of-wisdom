package server

// Storage interface
type Storage interface {
	GetRandomWisdom() string
}

// Service - server interface
type Service interface {
	ProcessMessage(message []byte, resource string) ([]byte, error)
}
