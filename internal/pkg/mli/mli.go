package mli

import "encoding/binary"

// Size - size of MLI header
const Size = 2

// EncodeWithMLI encodes the given message to message with MLI header.
func EncodeWithMLI(message []byte) []byte {
	b := make([]byte, Size)
	binary.BigEndian.PutUint16(b, uint16(len(message)+Size))

	return append(b, message...)
}

// GetMLI returns the MLI of the message.
func GetMLI(mli *[]byte) int {
	return int(binary.BigEndian.Uint16(*mli))
}
