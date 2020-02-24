package p02

import (
	"ctci/container/tree"
	"reflect"
	"testing"
)

func TestCreateMinimalBST(t *testing.T) {
	n := 7
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i + 1
	}
	type args struct {
		s []int
	}
	tr := &tree.Tree{Value: 4}
	tr.Insert(2)
	tr.Insert(6)
	tr.Insert(1)
	tr.Insert(3)
	tr.Insert(5)
	tr.Insert(7)

	tests := []struct {
		name string
		args args
		want *tree.Tree
	}{
		{"case1", args{s}, tr},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateMinimalBST(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeBTFromSortedSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
