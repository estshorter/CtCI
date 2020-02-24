package p04

import (
	"ctci/container/tree"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	tr := &tree.Tree{Value: 4}
	tr.Insert(2)
	tr.Insert(6)
	tr.Insert(1)
	tr.Insert(3)
	tr.Insert(7)

	tr2 := &tree.Tree{Value: 0}
	for i := 0; i < 6; i++ {
		tr2.Insert(i + 1)
	}

	type args struct {
		root *tree.Tree
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
			if got := IsBalanced(tt.args.root); got != tt.want {
				t.Errorf("IsBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
