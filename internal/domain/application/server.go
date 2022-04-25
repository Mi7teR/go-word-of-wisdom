package application

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/mi7ter/go-word-of-wisdom/internal/pkg/mli"
	"github.com/mi7ter/go-word-of-wisdom/internal/wisdom/server"
)

// RunServer runs the server
func RunServer(port int, service server.Service) {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Panic(fmt.Errorf("failed to listen: %w", err))
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(fmt.Errorf("failed to close listener: %w", err))
		}
	}(listener)

	log.Println("Listening on port", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(fmt.Errorf("failed to accept connection: %w", err))
			return
		}

		go handler(conn, service)
	}
}

func handler(conn net.Conn, service server.Service) {
	log.Println("New connection from address", conn.RemoteAddr())
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Panic(fmt.Errorf("failed to close connection: %w", err))
		}
	}(conn)

	for {
		b := make([]byte, mli.Size)

		_, err := conn.Read(b)
		if err != nil && err != io.EOF {
			log.Panic(fmt.Errorf("failed to read mli: %w", err))
		}

		if err == io.EOF {
			log.Println("Connection closed by client")
			return
		}

		length := mli.GetMLI(&b)
		mess := make([]byte, length)

		_, err = conn.Read(mess)
		if err != nil {
			log.Panic(fmt.Errorf("failed to read message: %w", err))
		}

		res, err := service.ProcessMessage(mess, strings.ReplaceAll(conn.RemoteAddr().String(), ":", ""))
		if err != nil {
			log.Println(fmt.Errorf("failed to process message: %w", err))

			err = conn.Close()
			if err != nil {
				log.Panic(fmt.Errorf("failed to close connection: %w", err))
			}

			log.Println("Connection closed")
		}

		_, err = conn.Write(mli.EncodeWithMLI(res))
		if err != nil {
			log.Panic(fmt.Errorf("failed to write response: %w", err))
		}
	}
}
