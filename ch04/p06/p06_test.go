package p06

import (
	"ctci/container/tree"
	"reflect"
	"testing"
)

func TestInOrderSucc(t *testing.T) {
	tr := &tree.Tree{Value: 4}
	tr.InsertWithParent(2)
	tr.InsertWithParent(6)
	tr.InsertWithParent(1)
	tr.InsertWithParent(3)
	tr.InsertWithParent(5)
	tr.InsertWithParent(7)

	type args struct {
		t *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want *tree.Tree
	}{
		{"case1", args{tr}, tr.Right.Left},
		{"case2", args{tr.Right.Right}, nil},
		{"case3", args{tr.Right.Left}, tr.Right},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InOrderSucc(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InOrderSucc() = %v, want %v", got, tt.want)
			}
		})
	}
}
