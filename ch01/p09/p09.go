package p09

import "strings"

// IsRotatedString checks if s2 is rotated matrix of s1
func IsRotatedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) || len(s1) == 0 || len(s2) == 0 {
		return false
	}
	return strings.Contains(s2+s2, s1)
}
