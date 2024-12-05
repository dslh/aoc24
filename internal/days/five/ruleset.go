package five

import "sort"

type rule struct {
	before int
	after  int
}

type ruleSet map[rule]struct{}

func (rs ruleSet) ordered(before int, after int) bool {
	_, ok := rs[rule{before, after}]
	return ok
}

func (rs ruleSet) Sort(update []int) {
	sort.Slice(update, func(i, j int) bool {
		return rs.ordered(update[i], update[j])
	})
}

func (rs ruleSet) Sorted(update []int) bool {
	return sort.SliceIsSorted(update, func(i, j int) bool {
		return rs.ordered(update[i], update[j])
	})
}

func makeRuleSet(rules []rule) ruleSet {
	ruleSet := make(ruleSet)
	for _, rule := range rules {
		ruleSet[rule] = struct{}{}
	}
	return ruleSet
}
