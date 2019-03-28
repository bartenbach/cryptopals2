package cryptopals2

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

// trim_bytes removes all newlines, carriage returns, and whitespace from byte arrays/slices.
// this function cannot apparently use a pointer receiver, which would be ideal.
func Delete_newlines(input []byte) []byte {
	// 10 is the ASCII value for a newline character, 13 is carriage return
	newline := []byte{ 10 }
	carriageReturn := []byte{ 13 }
	if bytes.Contains(input, newline) {
		input = append(input[:bytes.Index(input, newline)], input[bytes.Index(input, newline)+1:]...)
	}
	if bytes.Contains(input, carriageReturn) {
		input = append(input[:bytes.Index(input, carriageReturn)], input[bytes.Index(input, carriageReturn)+1:]...)
	}
	return input
}

// HammingDistance calculates the hamming distance between two byte arrays.
// This is the number of differing bits in binary representation.  The methodology implemented
// here was taken from Wikipedia where a similar algorithm was developed by Wegner (1960)
// https://en.wikipedia.org/wiki/Hamming_distance#Algorithm_example
func HammingDistance(input1 []byte, input2 []byte) (int, error) {
	if len(input1) != len(input2) {
		return -1, errors.New("byte arrays must be of equal length to calculate hamming distance")
	}
	/*
	totalDistance := 0
	for i := 0; i < len(input1); i++ {
		distance := input1[i] ^ input2[i]
		for distance > 2 {
		  totalDistance++
		  distance /= 2
		}
	}*/
	char1 := int8('a')
	char2 := int8('z')
	distance1 := char1 ^ char2
	fmt.Println(strconv.FormatInt(int64(char1), 2))
	fmt.Println(strconv.FormatInt(int64(char2), 2))
	// prints the visual representation of XOR output, all differing bits
	binaryXor := fmt.Sprintf("%07s", strconv.FormatInt(int64(char1 ^ char2), 2))
	fmt.Println(binaryXor)
	diffBits := 0
	for distance1 > 2 {
		diffBits++
		distance1 /= 2
	}
	fmt.Println("Total number of differing bits: ", diffBits)
	return diffBits, nil
}