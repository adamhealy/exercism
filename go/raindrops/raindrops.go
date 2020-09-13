// Package raindrops implements a function (Convert) to return raindrop sounds
package raindrops

import (
	"strconv"
)

// Convert converts a number into a string that contains raindrop sounds
// corresponding to certain potential factors.
func Convert(a int) string {
	s := ""
	if a%3 == 0 {
		s += "Pling"
	}
	if a%5 == 0 {
		s += "Plang"
	}
	if a%7 == 0 {
		s += "Plong"
	}
	if s == "" {
		s = strconv.Itoa(a)
	}
	return s
}
