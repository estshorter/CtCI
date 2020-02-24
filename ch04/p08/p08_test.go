package p08

import (
	"ctci/container/tree"
	"reflect"
	"testing"
)

func TestFindCommonAscendant(t *testing.T) {
	tr := &tree.Tree{Value: 4}
	tr.InsertWithParent(2)
	tr.InsertWithParent(6)
	tr.InsertWithParent(1)
	tr.InsertWithParent(3)
	tr.InsertWithParent(5)
	tr.InsertWithParent(7)
	type args struct {
		t1 *tree.Tree
		t2 *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want *tree.Tree
	}{
		{"case1", args{tr.Left, tr.Right}, tr},
		{"case2", args{tr.Right.Left, tr.Right.Right}, tr.Right},
		{"case3", args{tr, tr}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindCommonAscendant(tt.args.t1, tt.args.t2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindCommonAscendant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindCommonAscendant2(t *testing.T) {
	tr := &tree.Tree{Value: 4}
	tr.InsertWithParent(2)
	tr.InsertWithParent(6)
	tr.InsertWithParent(1)
	tr.InsertWithParent(3)
	tr.InsertWithParent(5)
	tr.InsertWithParent(7)

	type args struct {
		root *tree.Tree
		p    *tree.Tree
		q    *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want *tree.Tree
	}{
		{"case1", args{tr, tr.Left, tr.Right}, tr},
		{"case2", args{tr, tr.Right.Left, tr.Right.Right}, tr.Right},
		{"case3", args{tr, tr, tr}, tr},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindCommonAscendant2(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindCommonAscendant2() = %v, want %v", got, tt.want)
			}
		})
	}
}
