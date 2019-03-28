package cryptopals2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestSet1Challenge4(t *testing.T) {
	file, err := os.Open("test/4.txt")
	if err != nil {
		t.Fail()
	}
	corpus, err := BuildCorpus("test/A_Christmas_Carol.txt")
	if err != nil {
		t.Fail()
	}

	highestScore := -1
	bestCandidate := ""
	scanner := bufio.NewScanner(file)
	// get the line
	for scanner.Scan() {
		hexDecodedLine, err := DecodeHex(scanner.Text())
		if err != nil {
			t.Fail()
		}

		// score the line
		currentLineBest := BreakSingleCharXor(corpus, hexDecodedLine)

		// score the text
		score := ScoreText(corpus, []byte(currentLineBest))
		if score > highestScore {
			highestScore = score
			bestCandidate = string(currentLineBest)
		}
	}
	fmt.Println("Decrypted ciphertext is probably: ", bestCandidate)

	err = file.Close()
	if err != nil {
		log.Println("Failed to close challenge file 🤷")
	}
}
