package one

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

func part1(a []int, b []int) int {
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
	return total
}

func frequencies(a []int) map[int]int {
	freq := map[int]int{}
	for _, v := range a {
		freq[v]++
	}
	return freq
}

func part2(a []int, b []int) int {
	freq_b := frequencies(b)

	total := 0
	for _, v := range a {
		total += v * freq_b[v]
	}
	return total
}

func Main() {
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

	fmt.Println(part1(a, b))
	fmt.Println(part2(a, b))
}
