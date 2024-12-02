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

func unidirectional(diffs []int) bool {
	ascending := diffs[0] > 0
	for i := 1; i < len(diffs); i++ {
		ascends := diffs[i] > 0
		if ascending != ascends {
			return false
		}
	}
	return true
}

func gradual(diffs []int) bool {
	for _, val := range diffs {
		if val == 0 || val > 3 || val < -3 {
			return false
		}
	}
	return true
}

func safe(arr []int) bool {
	diffs := diffs(arr)
	return unidirectional(diffs) && gradual(diffs)
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

func Main() {
	reports, err := read_input()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(part1(reports))
}
