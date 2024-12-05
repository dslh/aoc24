package five

import (
	"bufio"
	"fmt"
	"os"
)

type rule struct {
	before int
	after  int
}

func updateOrder(update []int) map[int]int {
	order := make(map[int]int)
	for i, n := range update {
		order[n] = i
	}
	return order
}

func ruleMap(rules []rule) map[int][]rule {
	ruleMap := make(map[int][]rule)
	for _, rule := range rules {
		ruleMap[rule.before] = append(ruleMap[rule.before], rule)
	}
	return ruleMap
}

func correctOrder(updateOrder map[int]int, ruleMap map[int][]rule) bool {
	for page, index := range updateOrder {
		rules, exists := ruleMap[page]
		if !exists {
			continue
		}
		for _, rule := range rules {
			afterIndex, exists := updateOrder[rule.after]
			if !exists {
				continue
			}
			if afterIndex < index {
				return false
			}
		}
	}
	return true
}

func updateValue(update []int) int {
	i := len(update) / 2
	return update[i]
}

func part1(rules []rule, updates [][]int) int {
	total := 0
	ruleMap := ruleMap(rules)
	for _, update := range updates {
		updateOrder := updateOrder(update)
		if correctOrder(updateOrder, ruleMap) {
			total += updateValue(update)
		}
	}
	return total
}

func Main() {
	file, err := os.Open("input/5")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	rules, err := readRules(reader)
	if err != nil {
		fmt.Println("Error reading rules:", err)
		return
	}

	updates, err := readUpdates(reader)
	if err != nil {
		fmt.Println("Error reading updates:", err)
		return
	}

	fmt.Println(part1(rules, updates))
}
