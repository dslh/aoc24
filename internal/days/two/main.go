package two

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Reads the input into one array of integers per line.
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

// Returns an array of the differences between each pair of adjacent elements.
func diffs(arr []int) []int {
	diffs := []int{}
	for i := 1; i < len(arr); i++ {
		diffs = append(diffs, arr[i]-arr[i-1])
	}
	return diffs
}

// Returns true if the differences are all the same sign and false otherwise.
// Also returns the index of the first difference that is not the same sign.
// If all differences are the same sign, returns true and 0.
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

// Returns true if all differences are between -3 and 3 and false otherwise.
// Also returns the index of the first difference that is not between -3 and 3.
// Also returns false and the index of any element that is 0.
// If all differences are between -3 and 3, returns true and 0.
func gradual(diffs []int) (bool, int) {
	for i, val := range diffs {
		if val == 0 || val > 3 || val < -3 {
			return false, i
		}
	}
	return true, 0
}

// A report is "safe" if it is unidirectional and gradual.
func safe(report []int) bool {
	diffs := diffs(report)

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

// Returns a new array with the element at index i removed.
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

// Safe, but tolerant of a single bad level.
func tolerant_safe(arr []int) bool {
	diffs := diffs(arr)

	unidirectional, i := unidirectional(diffs)
	// We should be able to fix a fault by removing the level where the fault was found,
	// or the level before or after it.
	if !unidirectional {
		// safe_without also checks for gradual() faults. If we can't fix this fault,
		// there's no point trying to fix other faults.
		return safe_without(arr, i-1) || safe_without(arr, i) || safe_without(arr, i+1)
	}

	// We also need to fix faults where the report is unidirectional but not gradual.
	gradual, i := gradual(diffs[i:])
	if !gradual {
		return safe_without(arr, i-1) || safe_without(arr, i) || safe_without(arr, i+1)
	}

	// If we've made it this far, the report is safe.
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
