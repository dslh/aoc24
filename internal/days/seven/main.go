package seven

import "fmt"

func valid(c *calibration, ops []op) bool {
	for result := range permutations(c.equation, ops) {
		if result == c.result {
			return true
		}
	}
	return false
}

func totalValid(calibrations []*calibration, ops []op) int {
	total := 0
	for _, c := range calibrations {
		if valid(c, ops) {
			total += c.result
		}
	}
	return total
}

func Main() {
	calibrations, err := readInput()
	if err != nil {
		panic(err)
	}

	fmt.Println(totalValid(calibrations, []op{add, mul}))
	fmt.Println(totalValid(calibrations, []op{concat, add, mul}))
}
