package p05

func minProduct(a, b int) int {
	var bigger, smaller int
	if a > b {
		bigger = a
		smaller = b
	} else {
		bigger = b
		smaller = a
	}
	return minProductHelper(smaller, bigger)
}

func minProductHelper(smaller, bigger int) int {
	if smaller == 0 {
		return 0
	} else if smaller == 1 {
		return bigger
	}
	smallerHalf := smaller >> 1
	half := minProductHelper(smallerHalf, bigger)
	if smaller%2 == 0 {
		return half + half
	}

	return half + half + bigger
}
