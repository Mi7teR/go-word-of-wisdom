package storage

import (
	"bufio"
	"io"
	"math/rand"
	"os"
	"time"
)

// WisdomFileStorage is a server.Storage implementation that load wisdom list from a file and store them in memory.
type WisdomFileStorage struct {
	wisdomList  []string
	wisdomCount int
}

// NewWisdomStorage creates a new WisdomFileStorage instance.
func NewWisdomStorage(filePath string) (*WisdomFileStorage, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close() //nolint:errcheck

	wisdomList, err := LoadWisdomList(file)
	if err != nil {
		return nil, err
	}

	return &WisdomFileStorage{
		wisdomList:  wisdomList,
		wisdomCount: len(wisdomList),
	}, nil
}

// GetRandomWisdom returns random wisdom from the wisdom list.
func (w *WisdomFileStorage) GetRandomWisdom() string {
	if w.wisdomCount == 0 {
		return "when you trying to get wisdom from empty wisdom list, you will get this message"
	}

	rand.Seed(time.Now().Unix())

	return w.wisdomList[rand.Intn(w.wisdomCount)]
}

// LoadWisdomList loads wisdom list from an io.Reader interface.
func LoadWisdomList(reader io.Reader) ([]string, error) {
	var wisdomList []string

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		wisdomList = append(wisdomList, scanner.Text())
	}

	return wisdomList, nil
}
