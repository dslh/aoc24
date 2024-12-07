package common

import (
	"bufio"
	"unicode"
)

// Parse an integer from the reader.
// Returns the integer and the next rune, or an error.
func ReadInt(reader *bufio.Reader) (int, rune, error) {
	var n int
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			return n, 0, err
		}
		if unicode.IsDigit(r) {
			n = n*10 + int(r-'0')
		} else {
			return n, r, nil
		}
	}
}
