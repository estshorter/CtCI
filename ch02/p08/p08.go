package p08

import "ctci/container/sllist"

func findBeginning(l *sllist.List) *sllist.Element {
	eSlow := l.Front()
	eFast := l.Front()
	for eFast != nil && eFast.Next != nil {
		eSlow = eSlow.Next
		eFast = eFast.Next.Next
		if eSlow == eFast {
			break
		}
	}
	if eFast == nil || eFast.Next == nil {
		return nil
	}
	eSlow = l.Front()
	for eSlow != eFast {
		eSlow = eSlow.Next
		eFast = eFast.Next
	}
	return eSlow

}
