// Package hamming implements a function to return the Hamming Distance
// between two DNA strands
package hamming

import (
	"errors"
)

// Distance calculates the hamming distance given to DNA sequences
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)

	if len(ar) != len(br) {

		return 0, errors.New("DNA strands are not the same length")

	}

	sum := 0

	for i := 0; i < len(ar); i++ {

		if ar[i] != br[i] {
			sum++
		}

	}

	return sum, nil

}
