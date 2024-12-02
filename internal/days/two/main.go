package two

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_input() ([][]int, error) {
	file, err := os.Open("input/2")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reports := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Fields(text)
		report := []int{}
		for _, part := range parts {
			i, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			report = append(report, i)
		}
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func diffs(arr []int) []int {
	diffs := []int{}
	for i := 1; i < len(arr); i++ {
		diffs = append(diffs, arr[i]-arr[i-1])
	}
	return diffs
}

func unidirectional(diffs []int) (bool, int) {
	ascending := diffs[0] > 0
	for i := 1; i < len(diffs); i++ {
		ascends := diffs[i] > 0
		if ascending != ascends {
			return false, i
		}
	}
	return true, 0
}

func gradual(diffs []int) (bool, int) {
	for i, val := range diffs {
		if val == 0 || val > 3 || val < -3 {
			return false, i
		}
	}
	return true, 0
}

func safe(arr []int) bool {
	diffs := diffs(arr)

	unidirectional, _ := unidirectional(diffs)
	if !unidirectional {
		return false
	}

	gradual, _ := gradual(diffs)
	return gradual
}

func part1(reports [][]int) int {
	total := 0
	for _, report := range reports {
		if safe(report) {
			total++
		}
	}
	return total
}

func without(arr []int, i int) []int {
	result := make([]int, 0, len(arr)-1)
	result = append(result, arr[:i]...)
	result = append(result, arr[i+1:]...)
	return result
}

func safe_without(arr []int, i int) bool {
	if i < 0 || i >= len(arr) {
		return false
	}

	return safe(without(arr, i))
}

func tolerant_safe(arr []int) bool {
	diffs := diffs(arr)

	unidirectional, i := unidirectional(diffs)
	if !unidirectional {
		return safe_without(arr, i-1) || safe_without(arr, i) || safe_without(arr, i+1)
	}

	gradual, i := gradual(diffs[i:])
	if !gradual {
		return safe_without(arr, i-1) || safe_without(arr, i) || safe_without(arr, i+1)
	}

	return true
}

func part2(reports [][]int) int {
	total := 0
	for _, report := range reports {
		if tolerant_safe(report) {
			total++
		}
	}
	return total
}

func Main() {
	reports, err := read_input()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(part1(reports))
	fmt.Println(part2(reports))
}
