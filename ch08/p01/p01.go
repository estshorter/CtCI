package p01

func countStepNaive(n int) int {
	switch {
	case n < 0:
		return 0
	case n == 0:
		return 1
	default:
		return countStepNaive(n-1) + countStepNaive(n-2) + countStepNaive(n-3)
	}
}

func countStepMemo(n int) int {
	memo := make(map[int]int)
	return countStepMemoRecursive(n, memo)
}

func countStepMemoRecursive(n int, memo map[int]int) int {
	switch {
	case n < 0:
		return 0
	case n == 0:
		return 1
	default:
		a, ok := memo[n-1]
		if !ok {
			a = countStepMemoRecursive(n-1, memo)
			memo[n-1] = a
		}
		b, ok := memo[n-2]
		if !ok {
			b = countStepMemoRecursive(n-2, memo)
			memo[n-2] = b
		}
		c, ok := memo[n-3]
		if !ok {
			c = countStepMemoRecursive(n-3, memo)
			memo[n-3] = c
		}
		return a + b + c
	}
}
