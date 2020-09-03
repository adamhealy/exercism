// Package twofer returns one for you, and one for me.
package twofer

import "fmt"

// ShareWith shares with you or the name provided
func ShareWith(name string) string {
	// check if name exists, if not set to "you"
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
