package p01

import (
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortRune(s string) []rune {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return r
}

// NoDupMap checks if str has duplicted runes
func NoDupMap(str string) bool {
	m := make(map[rune]int)
	for _, runeValue := range str {
		m[runeValue]++
	}
	for _, cnt := range m {
		if cnt > 1 {
			return false
		}
	}
	return true
}

// NoDupLessMem checks if str has duplicted runes without map
func NoDupLessMem(str string) bool {
	runes := sortRune(str)
	for i := 1; i < len(runes); i++ {
		if runes[i-1] == runes[i] {
			return false
		}
	}
	return true
}
