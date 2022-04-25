package hasher

// Hasher is a combined interface for hashing functions.
type Hasher interface {
	Hash(data string) (string, error)
}
