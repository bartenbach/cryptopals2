package cryptopals2

import (
	"encoding/base64"
	"encoding/hex"
)

// HexToBase64 converts a Hex-encoded string to a base64 encoded string
func HexToBase64(input string) (string, error) {
	hexDecoded := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(hexDecoded, []byte(input))
	if err != nil {
		return "", err
	}
	base64encoded := make([]byte, base64.StdEncoding.EncodedLen(len(hexDecoded)))
	base64.StdEncoding.Encode(base64encoded, hexDecoded)

	return string(base64encoded), nil
}

// DecodeHex handles decoding of hex encoded strings.  Returns a byte array that contains the decoded result.
func DecodeHex(input string) ([]byte, error) {
	hexDecoded := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(hexDecoded, []byte(input))
	if err != nil {
		return nil, err
	}
	return hexDecoded, nil
}