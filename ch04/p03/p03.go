package p03

import (
	"container/list"
	"ctci/container/tree"
	"fmt"
)

// CreateLevelLinkedListBFS solves ch04-p03
func CreateLevelLinkedListBFS(root *tree.Tree) []*list.List {
	var result []*list.List
	current := list.New()
	current.PushBack(root)
	for current.Len() != 0 {
		result = append(result, current)
		parent := current
		current = list.New()
		for e := parent.Front(); e != nil; e = e.Next() {
			tr := e.Value.(*tree.Tree)
			if tr.Left != nil {
				current.PushBack(tr.Left)
			}
			if tr.Right != nil {
				current.PushBack(tr.Right)
			}
		}
	}
	return result
}

// CreateLevelLinkedListDFS solves ch04-p03
func CreateLevelLinkedListDFS(root *tree.Tree) []*list.List {
	var result []*list.List
	result = createLevelLinkedListDFSRecusive(root, 0, result)
	for _, v := range result {
		for e := v.Front(); e != nil; e = e.Next() {
			fmt.Printf("%v ", e.Value.(*tree.Tree).Value)
		}
		fmt.Println()
	}

	return result
}

func createLevelLinkedListDFSRecusive(tr *tree.Tree, depth int, result []*list.List) []*list.List {
	if tr == nil {
		return result
	}

	if len(result) == depth {
		l := list.New()
		l.PushBack(tr)
		result = append(result, l)
	} else {
		result[depth].PushBack(tr)
	}
	result = createLevelLinkedListDFSRecusive(tr.Left, depth+1, result)
	result = createLevelLinkedListDFSRecusive(tr.Right, depth+1, result)
	return result
}
