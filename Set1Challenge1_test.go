package cryptopals2

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSet1Challenge1(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	output := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	decoded, err := hex.DecodeString(input)
	if err != nil {
		t.Fail()
	}
	fmt.Println(string(decoded)) // print easter egg
	result, err := HexToBase64(input)
	if result != output || err != nil {
		t.Fail()
	}
}
