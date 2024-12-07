package seven

type op func(int, int) int

func concat(a, b int) int {
	for c := b; c > 0; c /= 10 {
		a *= 10
	}
	return a + b
}
func add(a, b int) int {
	return a + b
}
func mul(a, b int) int {
	return a * b
}

func permute(equation []int, ops []op, result int, i int, c chan int) {
	if i == len(equation) {
		c <- result
		return
	}

	for _, op := range ops {
		permute(equation, ops, op(result, equation[i]), i+1, c)
	}
}

func permutations(equation []int, ops []op) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		permute(equation, ops, equation[0], 1, c)
	}()
	return c
}
