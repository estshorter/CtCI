package p02

import "sort"

func anagramGroupingSort(a []string) []string {
	m := make(map[string][]string)
	for _, s := range a {
		b := []rune(s)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		if strSlice, ok := m[string(b)]; ok {
			m[string(b)] = append(strSlice, s)
		} else {
			m[string(b)] = []string{s}
		}
	}
	i := 0
	for _, strSlice := range m {
		for _, str := range strSlice {
			a[i] = str
			i++
		}
	}
	return a
}
