package cryptopals2

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSet1Challenge2(t *testing.T) {
	hexDecoded1, err := DecodeHex("1c0111001f010100061a024b53535009181c")
	hexDecoded2, err2 := DecodeHex("686974207468652062756c6c277320657965")
	expectedOutput, err3 := DecodeHex("746865206b696420646f6e277420706c6179")

	if err != nil || err2 != nil || err3 != nil {
		t.Fail()
	}
	// print easter eggs
	fmt.Println("Decoded1: ", string(hexDecoded1))
	fmt.Println("Decoded2: ", string(hexDecoded2))
	fmt.Println("Expected: ", string(expectedOutput))

	result, err := XorByteArrays(hexDecoded1, hexDecoded2)
	if err != nil {
		t.Fail()
	}
	if !bytes.Equal(result, expectedOutput) {
		t.Fail()
	}
}
