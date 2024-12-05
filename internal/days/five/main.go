package five

import (
	"fmt"
)

func updateValue(update []int) int {
	i := len(update) / 2
	return update[i]
}

func Main() {
	rules, updates, err := readInput()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	part1 := 0
	part2 := 0
	ruleSet := makeRuleSet(rules)
	for _, update := range updates {
		if ruleSet.Sorted(update) {
			part1 += updateValue(update)
		} else {
			ruleSet.Sort(update)
			part2 += updateValue(update)
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}
