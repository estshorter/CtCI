package p09

import (
	"ctci/container/list"
	"ctci/container/tree"
	"reflect"
	"testing"
)

func TestAllSequences(t *testing.T) {
	tr := &tree.Tree{Value: 10}
	tr.InsertWithParent(3)
	tr.InsertWithParent(2)
	tr.InsertWithParent(8)
	tr.InsertWithParent(9)

	results := make([]*list.List, 3)
	results[0] = list.New()
	results[0].PushBack(10)
	results[0].PushBack(3)
	results[0].PushBack(2)
	results[0].PushBack(8)
	results[0].PushBack(9)

	results[1] = list.New()
	results[1].PushBack(10)
	results[1].PushBack(3)
	results[1].PushBack(8)
	results[1].PushBack(2)
	results[1].PushBack(9)

	results[2] = list.New()
	results[2].PushBack(10)
	results[2].PushBack(3)
	results[2].PushBack(8)
	results[2].PushBack(9)
	results[2].PushBack(2)

	type args struct {
		t *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want []*list.List
	}{
		{"case1", args{tr}, results},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllSequences(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
