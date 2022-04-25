package application

import (
	"log"
	"net"
	"time"

	"github.com/mi7ter/go-word-of-wisdom/internal/wisdom/client"
)

// RunClient runs the client application
func RunClient(address string, service client.Service) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Panicf("Failed to connect to server: %v", err)
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Panicf("Failed to close connection: %v", err)
		}
	}(conn)

	log.Println("Connected to server")

	for {
		message, err := service.HandleConnection(conn)
		if err != nil {
			log.Panicf("Failed to handle connection: %v", err)
		}

		log.Printf("Received message: %s", message)

		time.Sleep(10 * time.Second)
	}
}
