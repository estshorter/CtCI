package p07

import "ctci/container/sllist"

func detectCommonNode(l1, l2 *sllist.List) []*sllist.Element {
	var ret []*sllist.Element
	lL, lS := l1, l2
	eL, eS := lL.Front(), lS.Front()
	if l1.Len() != l2.Len() {
		lL, lS = long(l1, l2)
		eL = lL.Front()
		for i := 0; i < lL.Len()-lS.Len(); i++ {
			eL = eL.Next
		}
	}
	for ; eS != nil; eL, eS = eL.Next, eS.Next {
		if eL == eS {
			ret = append(ret, eL)
		}
	}
	return ret
}

func long(x, y *sllist.List) (*sllist.List, *sllist.List) {
	if x.Len() < y.Len() {
		return y, x
	}
	return x, y
}
