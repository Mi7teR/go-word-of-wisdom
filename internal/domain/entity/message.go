package entity

import (
	"bytes"
	"encoding/gob"
)

// MessageType is a type of message
type MessageType int

// Message types
const (
	CloseMessageType MessageType = iota
	RequestChallengeMessageType
	ResponseChallengeMessageType
	WisdomMessageType
)

// Message is a message that can be sent between client and server.
type Message struct {
	Type    MessageType
	Payload []byte
}

// NewMessage creates a new message.
func NewMessage(t MessageType, payload []byte) *Message {
	return &Message{
		Type:    t,
		Payload: payload,
	}
}

// ToBytes implements the encoding.BinaryMarshaler interface.
func (m *Message) ToBytes() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// FromBytes creates a new message from bytes.
func FromBytes(b []byte) (*Message, error) {
	var mess Message

	dec := gob.NewDecoder(bytes.NewReader(b))

	if err := dec.Decode(&mess); err != nil {
		return nil, err
	}

	return &mess, nil
}

// String returns a string representation of the message.
func (m *Message) String() string {
	return string(m.Payload)
}
