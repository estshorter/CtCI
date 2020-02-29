package p02

import (
	"strings"
)

// PrintBinary solves ch05-p01
func PrintBinary(v float64) string {
	if v >= 1.0 || v <= 0 {
		return "ERROR"
	}
	sb := strings.Builder{}
	frac := 0.5
	sb.WriteRune('.')
	for v > 0 {
		if sb.Len() > 32 {
			return "ERROR"
		}

		if v >= frac {
			sb.WriteRune('1')
			v -= frac
		} else {
			sb.WriteRune('0')
		}
		frac /= 2
	}
	return sb.String()
}
