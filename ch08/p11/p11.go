package p11

func makeChange(n int) int {
	denoms := []int{25, 10, 5, 1}
	memo := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = make([]int, len(denoms))
	}
	return makeChangeHelper(n, denoms, 0, memo)
}

func makeChangeHelper(amount int, denoms []int, index int, memo [][]int) int {
	if memo[amount][index] > 0 {
		return memo[amount][index]
	}
	if index >= len(denoms)-1 {
		return 1
	}

	denomAmount := denoms[index]
	coinNumMax := amount / denomAmount
	ways := 0
	for i := 0; i <= coinNumMax; i++ {
		ways += makeChangeHelper(amount-i*denomAmount, denoms, index+1, memo)
	}
	memo[amount][index] = ways
	return ways
}
