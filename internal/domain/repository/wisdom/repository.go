package wisdom

import "github.com/mi7ter/go-word-of-wisdom/internal/wisdom/server"

// Repository - implementation of server.WisdomRepository
type Repository struct {
	storage server.Storage
}

// NewRepository - creates new wisdom repository
func NewRepository(storage server.Storage) *Repository {
	return &Repository{storage: storage}
}

// GetWisdom - get wisdom from repository
func (r Repository) GetWisdom() string {
	return r.storage.GetRandomWisdom()
}
