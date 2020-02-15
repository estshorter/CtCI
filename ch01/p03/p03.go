package p03

import "strings"

// UrlifyTest replace space with "%20"
func UrlifyTest(s string) string {
	rs := []rune(strings.TrimSpace(s))
	runes := make([]rune, 3*len(rs))
	for i, r := range rs {
		runes[i] = r
	}
	return string(urlify(runes, len(rs)))
}

func countSpace(rs []rune, n int) int {
	cnt := 0
	for i := 0; i < n; i++ {
		if rs[i] == ' ' {
			cnt++
		}
	}
	return cnt
}

func urlify(rs []rune, n int) []rune {
	ns := countSpace(rs, n)
	idxURL := n + 2*ns - 1
	for idxOrig := n - 1; idxOrig >= 0; idxOrig-- {
		if rs[idxOrig] == ' ' {
			rs[idxURL-2] = '%'
			rs[idxURL-1] = '2'
			rs[idxURL] = '0'
			idxURL -= 3
		} else {
			rs[idxURL] = rs[idxOrig]
			idxURL--
		}
	}
	return rs[:n+2*ns]
}
