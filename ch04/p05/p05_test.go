package p05

import (
	"ctci/container/tree"
	"testing"
)

func Test_isBST(t *testing.T) {
	tr := &tree.Tree{Value: 4}
	tr.Insert(2)
	tr.Insert(6)
	tr.Insert(1)
	tr.Insert(3)
	tr.Insert(7)

	tr2 := &tree.Tree{Value: 4}
	tr2.Insert(2)
	tr2.Insert(6)
	tr2.Insert(1)
	tr2.Insert(3)
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
