package main

import (
	"flag"
	"log"
	"time"

	"github.com/mi7ter/go-word-of-wisdom/internal/domain/application"
	"github.com/mi7ter/go-word-of-wisdom/internal/domain/repository/hashcashpow"
	"github.com/mi7ter/go-word-of-wisdom/internal/domain/service/client"
	"github.com/mi7ter/go-word-of-wisdom/internal/pkg/pow/hasher"
)

var (
	address       string
	maxIterations int
)

func main() {
	flag.StringVar(&address, "addr", "localhost:1337", "remote server address")
	flag.IntVar(&maxIterations, "i", 10000000, "max hasher iterations")
	flag.Parse()

	pow := hashcashpow.NewHashcashPow(maxIterations, hasher.NewSHA256(), 1*time.Hour)

	service := client.NewService(pow)

	application.RunClient(address, service)

	log.Println("client stopped working")
}
