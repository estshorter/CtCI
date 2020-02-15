package p04

func checkPalindromePermutation(s string) bool {
	rs := []rune(s)
	rCounter := make(map[rune]int)
	oddCounter := 0
	for _, r := range rs {
		if r != ' ' {
			rCounter[r]++
			val := rCounter[r]
			if val%2 != 0 {
				oddCounter++
			} else {
				oddCounter--
			}
		}
	}
	return oddCounter <= 1
}
