package three

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var mulRegex = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func readInput() (string, error) {
	bytes, err := os.ReadFile("input/3")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parseMul(mul []string) (int, int, error) {
	if len(mul) != 3 {
		return 0, 0, fmt.Errorf("expected 3 parts, got %d", len(mul))
	}
	a, err := strconv.Atoi(mul[1])
	if err != nil {
		return 0, 0, err
	}
	b, err := strconv.Atoi(mul[2])
	if err != nil {
		return 0, 0, err
	}
	return a, b, nil
}

func Main() {
	program, err := readInput()
	if err != nil {
		panic(err)
	}

	total := 0
	for _, match := range mulRegex.FindAllStringSubmatch(program, -1) {
		a, b, err := parseMul(match)
		if err != nil {
			panic(err)
		}
		total += a * b
	}
	fmt.Println(total)
}
