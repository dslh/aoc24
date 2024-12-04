package two

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Reads the input into one array of integers per line.
func readInput() ([][]int, error) {
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
func calculateDiffs(arr []int) []int {
	differences := []int{}
	for i := 1; i < len(arr); i++ {
		differences = append(differences, arr[i]-arr[i-1])
	}
	return differences
}

// Returns true if the differences are both unidirectional and gradual.
// Returns false and the index where either condition fails.
// Unidirectional means all differences have the same sign.
// Gradual means all differences are between -3 and 3 (exclusive of 0).
func checkDiffs(differences []int) (bool, int) {
	if len(differences) == 0 {
		return true, 0
	}

	ascending := differences[0] > 0

	for i, val := range differences {
		// Check gradual condition
		if val == 0 || val > 3 || val < -3 {
			return false, i
		}

		// Check unidirectional condition
		if i > 0 && (val > 0) != ascending {
			return false, i
		}
	}
	return true, 0
}

// A report is "safe" if its differences are unidirectional and gradual.
func isSafe(report []int) bool {
	differences := calculateDiffs(report)
	ok, _ := checkDiffs(differences)
	return ok
}

func part1(reports [][]int) int {
	total := 0
	for _, report := range reports {
		if isSafe(report) {
			total++
		}
	}
	return total
}

// Returns a new array with the element at index i removed.
func removeElement(arr []int, i int) []int {
	result := make([]int, 0, len(arr)-1)
	result = append(result, arr[:i]...)
	result = append(result, arr[i+1:]...)
	return result
}

func isSafeWithout(arr []int, i int) bool {
	if i < 0 || i >= len(arr) {
		return false
	}

	return isSafe(removeElement(arr, i))
}

// Safe, but tolerant of a single bad level.
func isTolerantSafe(arr []int) bool {
	differences := calculateDiffs(arr)

	ok, i := checkDiffs(differences)
	if !ok {
		return isSafeWithout(arr, i-1) || isSafeWithout(arr, i) || isSafeWithout(arr, i+1)
	}

	return true
}

func part2(reports [][]int) int {
	total := 0
	for _, report := range reports {
		if isTolerantSafe(report) {
			total++
		}
	}
	return total
}

func Main() {
	reports, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(part1(reports))
	fmt.Println(part2(reports))
}
