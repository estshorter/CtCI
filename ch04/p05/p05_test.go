package p05

import (
	"ctci/container/tree"
	"testing"
)

func Test_isBST(t *testing.T) {
	tr := &tree.Tree{Value: 4}
	tr = tree.Insert(tr, 2)
	tr = tree.Insert(tr, 6)
	tr = tree.Insert(tr, 1)
	tr = tree.Insert(tr, 3)
	// tr = tree.Insert(tr, 5)
	tr = tree.Insert(tr, 7)

	tr2 := &tree.Tree{Value: 4}
	tr2 = tree.Insert(tr2, 2)
	tr2 = tree.Insert(tr2, 6)
	tr2 = tree.Insert(tr2, 1)
	tr2 = tree.Insert(tr2, 3)
	tr3 := &tree.Tree{Value: 7}
	tr2.Right.Left = tr3

	type args struct {
		t *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{tr}, true},
		{"case2", args{tr2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBST(tt.args.t); got != tt.want {
				t.Errorf("isBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
