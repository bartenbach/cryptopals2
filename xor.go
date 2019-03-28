package cryptopals2

import (
	"errors"
	"strconv"
)

// XorByteArrays performs an XOR on two byte arrays and returns the result
// in a byte array
func XorByteArrays(bytes1, bytes2 []byte) ([]byte, error) {
	if len(bytes1) != len(bytes2) {
		val1 := strconv.Itoa(len(bytes1))
		val2 := strconv.Itoa(len(bytes2))
		return nil, errors.New("byte buffers must be of equal length (" + val1 + " and " + val2 + ")")
	}
	xored := make([]byte, len(bytes1))
	for i := range bytes1 {
		xored[i] = bytes1[i] ^ bytes2[i]
	}
	return xored, nil
}

// SingleCharacterXor performs an XOR on an array of bytes using a single character.
func SingleCharacterXor(input []byte, key byte) []byte {
	xored := make([]byte, len(input))
	for i:= range input {
		xored[i] = input[i] ^ key
	}
	return xored
}

// Returns a repeating key XOR from the specified key of the specified length.
func CreateRepeatingKeyXOR(key []byte, length int) []byte {
	repeatedKey := make([]byte, length)
	j := 0
	for i := 0; i < length; i ++ {
		repeatedKey[i] = key[j]
		if j < len(key) - 1 {
			j++
		} else {
			j = 0
		}
	}
	return repeatedKey
}