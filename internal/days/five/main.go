package five

import (
	"fmt"
	"sort"
)

type rule struct {
	before int
	after  int
}

type ruleSet map[rule]struct{}

func (rs ruleSet) Ordered(before int, after int) bool {
	_, ok := rs[rule{before, after}]
	return ok
}
func (rs ruleSet) Sort(update []int) {
	sort.Slice(update, func(i, j int) bool {
		return rs.Ordered(update[i], update[j])
	})
}

func makeRuleSet(rules []rule) ruleSet {
	ruleSet := make(ruleSet)
	for _, rule := range rules {
		ruleSet[rule] = struct{}{}
	}
	return ruleSet
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

func correctOrder(update []int, ruleMap map[int][]rule) bool {
	updateOrder := updateOrder(update)
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

func Main() {
	rules, updates, err := readInput()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	part1 := 0
	part2 := 0
	ruleMap := ruleMap(rules)
	ruleSet := makeRuleSet(rules)
	for _, update := range updates {
		if correctOrder(update, ruleMap) {
			part1 += updateValue(update)
		} else {
			ruleSet.Sort(update)
			part2 += updateValue(update)
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}
