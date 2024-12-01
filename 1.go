package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func extract_lists(file *os.File) ([]int, []int, error) {
	a := []int{}
	b := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		parts := strings.Fields(text)
		i, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}
		a = append(a, i)

		j, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}
		b = append(b, j)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	return a, b, nil
}

func main() {
	file, err := os.Open("input/1")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	a, b, err := extract_lists(file)
	if err != nil {
		fmt.Println("Error extracting lists:", err)
		return
	}

	sort.Ints(a)
	sort.Ints(b)

	total := 0
	for i := 0; i < len(a); i++ {
		diff := a[i] - b[i]
		if diff < 0 {
			diff = -diff
		}
		total += diff
	}
	fmt.Println(total)
}
