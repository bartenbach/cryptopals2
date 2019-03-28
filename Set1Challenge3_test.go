package cryptopals2

import (
	"fmt"
	"testing"
)

func TestSet1Challenge3(t *testing.T) {
	encryptedString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	hexDecodedInput, _ := DecodeHex(encryptedString)
	corpus, err := BuildCorpus("test/A_Christmas_Carol.txt")
	if err != nil {
		t.Fail()
	}

	highestScore := -1
	bestCandidate := ""
	// iterate through all ASCII character values
	for i := 0; i < 128; i++ {
		xored := SingleCharacterXor(hexDecodedInput, byte(i))
		score := ScoreText(corpus, xored)
		if score > highestScore {
			highestScore = score
			bestCandidate = string(xored)
		}
	}
	fmt.Println("Decrypted ciphertext is probably: ", bestCandidate)
}
