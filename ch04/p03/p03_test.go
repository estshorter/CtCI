package p03

import (
	"container/list"
	"ctci/container/tree"
	"fmt"
	"reflect"
	"testing"
)

func TestCreateLevelLinkedListBFS(t *testing.T) {
	tr := &tree.Tree{Value: 4}
	tr.Insert(2)
	tr.Insert(6)
	tr.Insert(1)
	tr.Insert(3)
	tr.Insert(5)
	tr.Insert(7)

	var ls []*list.List
	l1 := list.New()
	l1.PushBack(tr)
	ls = append(ls, l1)
	l2 := list.New()
	l2.PushBack(tr.Left)
	l2.PushBack(tr.Right)
	ls = append(ls, l2)
	l3 := list.New()
	l3.PushBack(tr.Left.Left)
	l3.PushBack(tr.Left.Right)
	l3.PushBack(tr.Right.Left)
	l3.PushBack(tr.Right.Right)
	ls = append(ls, l3)

	type args struct {
		root *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want []*list.List
	}{
		{"case1", args{tr}, ls},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateLevelLinkedListBFS(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateLevelLinkedListBFS() = %v, want %v", got, tt.want)
			}
			for _, v := range tt.want {
				for e := v.Front(); e != nil; e = e.Next() {
					fmt.Printf("%v ", e.Value.(*tree.Tree).Value)
				}
				fmt.Println()
			}
		})
	}
}

func TestCreateLevelLinkedListDFS(t *testing.T) {
	tr := &tree.Tree{Value: 4}
	tr.Insert(2)
	tr.Insert(6)
	tr.Insert(1)
	tr.Insert(3)
	tr.Insert(5)
	tr.Insert(7)

	var ls []*list.List
	l1 := list.New()
	l1.PushBack(tr)
	ls = append(ls, l1)
	l2 := list.New()
	l2.PushBack(tr.Left)
	l2.PushBack(tr.Right)
	ls = append(ls, l2)
	l3 := list.New()
	l3.PushBack(tr.Left.Left)
	l3.PushBack(tr.Left.Right)
	l3.PushBack(tr.Right.Left)
	l3.PushBack(tr.Right.Right)
	ls = append(ls, l3)

	type args struct {
		root *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want []*list.List
	}{
		{"case1", args{tr}, ls},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateLevelLinkedListDFS(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateLevelLinkedListDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
