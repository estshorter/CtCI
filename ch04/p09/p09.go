package p09

import (
	"ctci/container/list"
	"ctci/container/tree"
)

// AllSequences solves ch04-p09
func AllSequences(t *tree.Tree) []*list.List {
	var results []*list.List
	if t == nil {
		l := list.New()
		results = append(results, l)
		return results
	}

	prefix := list.New()
	prefix.PushBack(t.Value)

	resL := AllSequences(t.Left)
	resR := AllSequences(t.Right)

	for _, lL := range resL {
		for _, lR := range resR {
			results = weaveList(lL, lR, prefix, results)
		}
	}
	return results
}

func weaveList(first, second, prefix *list.List, results []*list.List) []*list.List {
	if first.Len() == 0 || second.Len() == 0 {
		result := list.New()
		result.PushBackListSafe(prefix)
		result.PushBackListSafe(first)
		result.PushBackListSafe(second)
		return append(results, result)
	}

	headFirst := first.GetElement(0)
	first.Remove(headFirst)
	prefix.PushBack(headFirst.Value)
	results = weaveList(first, second, prefix, results)
	prefix.Remove(prefix.GetElement(prefix.Len() - 1))
	first.PushFront(headFirst.Value)

	headSecond := second.GetElement(0)
	second.Remove(headSecond)
	prefix.PushBack(headSecond.Value)
	results = weaveList(first, second, prefix, results)
	prefix.Remove(prefix.GetElement(prefix.Len() - 1))
	second.PushFront(headSecond.Value)

	return results
}
