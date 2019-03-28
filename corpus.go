package cryptopals2

import (
	"fmt"
	"io/ioutil"
)

// BuildCorpus builds a corpus of word frequency from the specified file.
// If the file cannot be opened, returns nil.
func BuildCorpus(corpusPath string) (map[string]int, error) {
	corpusText, err := ioutil.ReadFile(corpusPath)
	if err != nil {
		err := fmt.Errorf("failed to read %s", corpusPath)
		return nil, err
	}

	corpus := make(map[string]int)

    for _, s := range corpusText {
    	corpus[string(s)]++
	}

	return corpus, nil
}

func ScoreText(corpus map[string]int, text []byte) int {
	score := 0
	for _,s := range text {
		score += corpus[string(s)]
	}
	return score
}
