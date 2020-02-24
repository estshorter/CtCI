package p10

import (
	"ctci/container/tree"
	"strings"
	"testing"
)

func Test_getOrderString(t *testing.T) {
	tr := &tree.Tree{Value: 4}
	tr.InsertWithParent(2)
	tr.InsertWithParent(6)
	tr.InsertWithParent(1)
	sb := &strings.Builder{}
	type args struct {
		t  *tree.Tree
		sb *strings.Builder
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{tr, sb}, "4 2 1 X X X 6 X X "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getOrderString(tt.args.t, tt.args.sb)
			got := sb.String()
			if got != tt.want {
				t.Errorf("got :%v, want:%v", got, tt.want)
			}
		})
	}
}

func Test_containsTree(t *testing.T) {
	tr := &tree.Tree{Value: 4}
	tr.InsertWithParent(2)
	tr.InsertWithParent(6)
	tr.InsertWithParent(1)
	tr.InsertWithParent(7)
	tr.InsertWithParent(9)

	tr2 := &tree.Tree{Value: 2}
	tr2.InsertWithParent(4)

	type args struct {
		t1 *tree.Tree
		t2 *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{tr, tr.Right}, true},
		{"case2", args{tr, tr2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsTree(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("containsTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsTree2(t *testing.T) {
	tr := &tree.Tree{Value: 4}
	tr.InsertWithParent(2)
	tr.InsertWithParent(6)
	tr.InsertWithParent(1)
	tr.InsertWithParent(7)
	tr.InsertWithParent(9)

	tr2 := &tree.Tree{Value: 2}
	tr2.InsertWithParent(4)

	type args struct {
		t1 *tree.Tree
		t2 *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{tr, tr.Right}, true},
		{"case2", args{tr, tr2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsTree2(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("ContainsTree2() = %v, want %v", got, tt.want)
			}
		})
	}
}
