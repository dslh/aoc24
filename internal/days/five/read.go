package five

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"unicode"
)

var ErrEmptyLine = errors.New("empty line")

// Parse an integer from the reader.
// Returns the integer and the next rune, or an error.
func readInt(reader *bufio.Reader) (int, rune, error) {
	var n int
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			return 0, 0, err
		}
		if unicode.IsDigit(r) {
			n = n*10 + int(r-'0')
		} else {
			return n, r, nil
		}
	}
}

func readRule(reader *bufio.Reader) (rule, error) {
	before, separator, err := readInt(reader)
	if err != nil {
		return rule{}, err
	}
	if separator == '\n' && before == 0 {
		return rule{}, ErrEmptyLine
	}
	if separator != '|' {
		return rule{}, fmt.Errorf("expected '|', got %c", separator)
	}

	after, newline, err := readInt(reader)
	if err != nil {
		return rule{}, err
	}
	if newline != '\n' {
		return rule{}, fmt.Errorf("expected newline, got %c", newline)
	}

	return rule{before, after}, nil
}

func readRules(reader *bufio.Reader) ([]rule, error) {
	var rules []rule
	for {
		rule, err := readRule(reader)
		if err == ErrEmptyLine {
			break
		}
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}
	return rules, nil
}

func readUpdate(reader *bufio.Reader) ([]int, error) {
	var updates []int
	for {
		n, separator, err := readInt(reader)
		if err != nil {
			return nil, err
		}
		updates = append(updates, n)
		if separator == '\n' {
			break
		}
		if separator != ',' {
			return nil, fmt.Errorf("expected ',', got %c", separator)
		}
	}
	return updates, nil
}

func readUpdates(reader *bufio.Reader) ([][]int, error) {
	var updates [][]int
	for {
		update, err := readUpdate(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		updates = append(updates, update)
	}
	return updates, nil
}

func readInput() ([]rule, [][]int, error) {
	file, err := os.Open("input/5")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	rules, err := readRules(reader)
	if err != nil {
		return nil, nil, err
	}

	updates, err := readUpdates(reader)
	if err != nil {
		return nil, nil, err
	}

	return rules, updates, nil
}
