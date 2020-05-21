package p04

func subset(a []string) []string {
	if len(a) == 0 {
		return []string{""}
	}
	subsub := subset(a[0 : len(a)-1])
	ret := make([]string, 2*len(subsub))
	for i := 0; i < len(subsub); i++ {
		ret[i] = subsub[i]
		ret[i+len(subsub)] = subsub[i] + a[len(a)-1]
	}
	return ret
}
