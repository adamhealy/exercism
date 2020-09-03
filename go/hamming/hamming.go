// Package hamming implements a function to return the Hamming Distance
// between two DNA strands
package hamming

import (
	"errors"
	"strings"
)

// Distance calculates the hamming distance given to DNA sequences
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {

		return -1, errors.New("DNA strands are not the same length")

	}

	sliceA := strings.Split(a, "")
	sliceB := strings.Split(b, "")

	sum := 0

	for i := 0; i < len(a); i++ {

		if sliceA[i] != sliceB[i] {
			sum++
		}

	}

	return sum, nil

}
