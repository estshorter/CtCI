package p04

func subset(a []string) []string {
	if len(a) == 1 {
		return a
	}
	subsub := subset(a[0 : len(a)-1])
	ret := make([]string, 2*len(subsub)+1)
	for i := 0; i < len(subsub); i++ {
		ret[i] = subsub[i]
	}
	for i := 0; i < len(subsub); i++ {
		ret[i+len(subsub)] = subsub[i] + a[len(a)-1]
	}
	ret[len(ret)-1] = a[len(a)-1]
	return ret
}
