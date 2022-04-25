package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/mi7ter/go-word-of-wisdom/internal/domain/application"
	"github.com/mi7ter/go-word-of-wisdom/internal/domain/repository/hashcashpow"
	"github.com/mi7ter/go-word-of-wisdom/internal/domain/repository/wisdom"
	"github.com/mi7ter/go-word-of-wisdom/internal/domain/service/server"
	"github.com/mi7ter/go-word-of-wisdom/internal/pkg/pow/hasher"
	"github.com/mi7ter/go-word-of-wisdom/internal/pkg/storage"
)

var (
	wisdomFilePath string
	port           int
	bits           int
	maxIterations  int
)

func main() {
	flag.StringVar(&wisdomFilePath, "f", "wisdoms.txt", "filepath for wisdom file storage")
	flag.IntVar(&port, "p", 1337, "listening port")
	flag.IntVar(&bits, "b", 5, "zero bits length")
	flag.IntVar(&maxIterations, "i", 10000000, "max hasher iterations")
	flag.Parse()

	pow := hashcashpow.NewHashcashPow(maxIterations, hasher.NewSHA256(), 1*time.Hour)

	fileStorage, err := storage.NewWisdomStorage(wisdomFilePath)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create wisdom storage: %w", err))
	}

	repo := wisdom.NewRepository(fileStorage)

	service := server.NewService(pow, repo, bits)

	application.RunServer(port, service)

	log.Println("server stopped working")
}
