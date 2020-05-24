package p09

func addParen(list []string, leftRem, rightRem int, str []rune, index int) []string {
	if leftRem < 0 || rightRem < leftRem {
		return list
	} else if leftRem == 0 && rightRem == 0 {
		return append(list, string(str))
	}

	str[index] = '('
	list = addParen(list, leftRem-1, rightRem, str, index+1)

	str[index] = ')'
	list = addParen(list, leftRem, rightRem-1, str, index+1)
	return list
}

func generateParens(count int) []string {
	str := make([]rune, 6)
	list := []string{}
	return addParen(list, count, count, str, 0)
}
