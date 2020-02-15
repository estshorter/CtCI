package p05

// CanOneTimeConvert returns true if s1 can be s2 by one insertion/removal/replace or no-operation
func CanOneTimeConvert(s1 string, s2 string) bool {
	rs1 := []rune(s1)
	rs2 := []rune(s2)

	if len(rs1) == len(rs2)+1 {
		return EqualIfInsert(rs2, rs1)
	} else if len(s1)+1 == len(s2) {
		return EqualIfInsert(rs1, rs2)
	} else if len(rs1) == len(rs2) {
		return EqualIfReplaceOrNOOP(rs1, rs2)
	}
	return false
}

// EqualIfReplaceOrNOOP returns true if s1 can be s2 by one replace or no-operation
// assume len(a) = len(b)
func EqualIfReplaceOrNOOP(a []rune, b []rune) bool {
	foundDifference := false
	for idx, r := range a {
		if r != b[idx] {
			if foundDifference {
				return false
			}
			foundDifference = true
		}
	}
	return true
}

// EqualIfInsert returns true if s1 can be s2 by one insertion
// assume len(a) < len(b)
func EqualIfInsert(a []rune, b []rune) bool {
	ib := 0
	for ia := 0; ia < len(a); ia++ {
		if a[ia] != b[ib] {
			if ia != ib {
				return false
			}
			ib++
		}
		ib++
	}
	return true
}
