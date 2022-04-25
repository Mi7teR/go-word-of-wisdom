# go-word-of-wisdom - TCP-server with protection from DDOS attacks based on Proof of Work
This is a solution for interview test task

## Task
Design and implement "Word of Wisdom" tcp server.
- TCP server should be protected from DDOS attacks with the [Proof of Work] (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
- The choice of the POW algorithm should be explained.
- After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.
- Docker file should be provided both for the server and for the client that solves the POW challenge

## Chosen algorithm
The chosen algorithm is [Hashcash](https://en.wikipedia.org/wiki/Hashcash).

### Why Hashcash?
- It is a proof-of-work algorithm with good documentation.
- It is a simple algorithm with low computational complexity.
- It is a very popular algorithm.
- It is a very easy to implement algorithm.
- Validation of the proof of work is very simple.

## Protocol description
Client and server are communicating via TCP connection. The protocol use the following custom structure for messages:
```go
package entity

type MessageType int

const (
	CloseMessageType MessageType = iota
	RequestChallengeMessageType
	ResponseChallengeMessageType
	WisdomMessageType
)

type Message struct {
	Type    MessageType
	Payload []byte
}
```
Message is encoding to bytes before sending and decoding by [encoding/gob](https://pkg.go.dev/encoding/gob) after receiving.
Before encoded message sent, message length indicator(MLI) is added to the beginning of the message.
When message received, MLI is removed from the beginning of the message.
MLI is encoded to bytes using [binary.BigEndian](https://pkg.go.dev/encoding/binary?tab=doc#BigEndian),
using 2 bytes. MLI is used to determine the length of the message and read the full message from the connection.

### Protocol steps
- Client sends a message with the challenge request.
- Server sends a message with the response with included challenge.
- Client sends a message with the computed challenge.
- Server sends a message with the response with included random wisdom.

## How to run?
### Requirements
- Docker
- Docker-compose

### Commands

```shell
docker-compose up --abort-on-container-exit --force-recreate --build
```