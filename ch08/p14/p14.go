package p14

// Key defines key for map
type Key struct {
	result bool
	s      string
}

func countEval(s string, result bool, memo map[Key]int) int {
	if len(s) == 1 {
		if stringToBool(s) == result {
			return 1
		}
		return 0
	}

	if v, ok := memo[Key{result, s}]; ok {
		return v
	}

	ways := 0
	for i := 1; i < len(s); i += 2 {
		c := rune(s[i])
		left := s[0:i]
		right := s[i+1:]
		leftTrue := countEval(left, true, memo)
		leftFalse := countEval(left, false, memo)
		rightTrue := countEval(right, true, memo)
		rightFalse := countEval(right, false, memo)
		total := (leftTrue + leftFalse) * (rightTrue + rightFalse)

		totalTrue := 0
		if c == '^' {
			totalTrue = leftTrue*rightFalse + leftFalse*rightTrue
		} else if c == '&' {
			totalTrue = leftTrue * rightTrue
		} else if c == '|' {
			totalTrue = leftTrue*rightTrue + leftFalse*rightTrue + leftTrue*rightFalse
		}

		if result {
			ways += totalTrue
		} else {
			ways += total - totalTrue
		}

	}
	memo[Key{result, s}] = ways
	return ways
}

func stringToBool(c string) bool {
	return c == "1"
}
