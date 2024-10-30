package packets

import (
	"bytes"
	"fmt"
)

func Uint16be(data []byte) uint16 {
	return uint16(data[1]) | uint16(data[0]) << 8
}

func WriteUint16be(data []byte, value uint16) {
	data[0] = byte(value >> 8)
	data[1] = byte(value)
}

// LengthPrefixedString represents a string with a 16-bit length prefix.

type LengthPrefixedString struct {
	// Length of the string
	Length uint16
	// String data
	Value string
}

// Unmarshal decodes a length-prefixed string from the provided byte slice.
// It sets the Length and Value fields of the LengthPrefixedString receiver.
// The function returns the number of bytes read from the input data.
//
// Parameters:
//   - data: A byte slice containing the length-prefixed string to be unmarshaled.
//
// Returns:
//   - An integer representing the total number of bytes read from the input data.
func (lps *LengthPrefixedString) Unmarshal(data []byte) int {
	lps.Length = Uint16be(data)
	fmt.Println("Length: ", lps.Length)
	if len(data) < 2+int(lps.Length) {
		fmt.Println("Error: data slice too short for the specified length")
		return 0
	}
	lps.Value = string(data[2:2+lps.Length])
	fmt.Println("Data: ", lps.Value)
	return 2 + int(lps.Length)
}

func (lps *LengthPrefixedString) Marshal() []byte {
	var buffer bytes.Buffer
	WriteUint16be(buffer.Bytes(), lps.Length)
	buffer.Write([]byte(lps.Value))
	return buffer.Bytes()
}

func (lps *LengthPrefixedString) ToString() string {
	return string(lps.Value)
}