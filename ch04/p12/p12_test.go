package p12

import (
	"ctci/container/tree"
	"testing"
)

func Test_getSumPath(t *testing.T) {
	tr := &tree.Tree{Value: 4}
	tr.InsertWithParent(2)
	tr.InsertWithParent(6)
	tr.InsertWithParent(1)
	tr.InsertWithParent(3)
	tr.InsertWithParent(5)
	tr.InsertWithParent(7)

	type args struct {
		t      *tree.Tree
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{tr, 6}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSumPath(tt.args.t, tt.args.target); got != tt.want {
				t.Errorf("getSumPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountPathsWithSum(t *testing.T) {
	tr := &tree.Tree{Value: 4}
	tr.InsertWithParent(2)
	tr.InsertWithParent(6)
	tr.InsertWithParent(1)
	tr.InsertWithParent(3)
	tr.InsertWithParent(5)
	tr.InsertWithParent(7)

	type args struct {
		root      *tree.Tree
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{tr, 6}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountPathsWithSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("CountPathsWithSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
