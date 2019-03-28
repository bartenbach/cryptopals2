package cryptopals2

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSet1Challenge5(t *testing.T) {
	input := []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`)
	expectedOutput := []byte(`0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272
a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`)

	// create the key for the repeating key XOR
	repeatingKey := []byte("ICE")
	repeatedKey := CreateRepeatingKeyXOR(repeatingKey, len(input))

	// perform the XOR
	result, err := XorByteArrays(input, repeatedKey)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// hex encode the result
	hexEncodedResult := make([]byte, hex.EncodedLen(len(result)))
	hex.Encode(hexEncodedResult, result)

	// normalize the arrays
	expectedOutput = Delete_newlines(expectedOutput)
	trimmedHexEncodedResult := Delete_newlines(hexEncodedResult)

	// assert array equals expected array
	if !bytes.Equal(trimmedHexEncodedResult, expectedOutput) {
		fmt.Println("Found:    ", string(hexEncodedResult))
		fmt.Println("Expected: ", string(expectedOutput))
		t.Fail()
	}

}
