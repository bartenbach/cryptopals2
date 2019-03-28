package cryptopals2

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestSet1Challenge6(t *testing.T) {

	file, err := ioutil.ReadFile("test/6.txt")
	if err != nil {
		log.Fatal("Failed to open challenge file")
		t.Fail()
	}
	fmt.Println(string(file))

}

func TestHammingDistance(t *testing.T) {
	input1 := []byte("wokka wokka!!!")
	input2 := []byte("this is a test")
	distance, err := HammingDistance(input1, input2)
	if err != nil {
		log.Fatal(err)
		t.Fail()
	}
	if distance != 37 {
		log.Fatal("Calculated incorrect hamming distance")
		t.Fail()
	}
}