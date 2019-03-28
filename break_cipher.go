package cryptopals2

// BreakSingleCharXor accepts both a corpus and a byte array.  The function will iterate through
// all ASCII values and find the best scoring text based upon the corpus provided.  This text may or may
// not actually be english text, but it will be the "best guess".
func BreakSingleCharXor(corpus map[string]int, input []byte) string {
	highestScore := -1
	bestCandidate := ""
	// iterate through all ASCII character values
	for i := 0; i < 128; i++ {
		xored := SingleCharacterXor(input, byte(i))
		score := ScoreText(corpus, xored)
		if score > highestScore {
			highestScore = score
			bestCandidate = string(xored)
		}
	}
	return bestCandidate
}