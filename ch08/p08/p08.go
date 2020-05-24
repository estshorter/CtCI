package p08

func permutation(input string) []string {
	return permHelper("", len(input), buildRuneCount(input), []string{})
}

func permHelper(prefix string, reminder int, counter map[rune]int, result []string) []string {
	if reminder == 0 {
		return append(result, prefix)
	}
	for r, c := range counter {
		if c > 0 {
			counter[r]--
			result = permHelper(prefix+string(r), reminder-1, counter, result)
			counter[r]++
		}
	}
	return result
}

func buildRuneCount(input string) map[rune]int {
	counter := make(map[rune]int)
	for _, r := range input {
		counter[r]++
	}
	return counter
}
