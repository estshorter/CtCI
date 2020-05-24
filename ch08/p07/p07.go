package p07

func permutation(input string) []string {
	if len(input) == 1 {
		return []string{input}
	}
	firstRune := rune(input[0])
	sub := permutation(input[1:])
	ret := []string{}
	for _, str := range sub {
		for i := 0; i < len(str)+1; i++ {
			ret = append(ret, str[:i]+string(firstRune)+str[i:])
		}
	}
	return ret
}
