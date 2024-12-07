package seven

import "fmt"

func valid(c *calibration) bool {
	permutations := 1 << (len(c.equation) - 1)
	for i := 0; i < permutations; i++ {
		sum := c.equation[0]
		for j := 0; j < len(c.equation)-1; j++ {
			if i&(1<<j) != 0 {
				sum += c.equation[j+1]
			} else {
				sum *= c.equation[j+1]
			}
		}
		if sum == c.result {
			return true
		}
	}
	return false
}

func totalValid(calibrations []*calibration) int {
	total := 0
	for _, c := range calibrations {
		if valid(c) {
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

	fmt.Println(totalValid(calibrations))
}
