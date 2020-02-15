package p02

func makeRuneMap(s string) map[rune]int {
	s1Rune := []rune(s)
	sMap := make(map[rune]int)
	for _, r := range s1Rune {
		sMap[r]++
	}
	return sMap
}

// CheckPermutation checks s1 is permuation of s2
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1Map := makeRuneMap(s1)
	s2Rune := []rune(s2)
	for _, r2 := range s2Rune {
		if cnt, ok := s1Map[r2]; ok == false {
			return false
		} else if cnt == 0 {
			return false
		} else {
			s1Map[r2]--
		}
	}
	return true
}

// func main() {
// 	fmt.Print(checkPermutation("avv", "vva"))
// }
