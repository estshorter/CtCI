package tree

import (
	"testing"
)

func TestTree_InsertWithParent(t *testing.T) {
	tr := &Tree{Value: 4}
	tr.InsertWithParent(2)
	tr.InsertWithParent(6)
	tr.InsertWithParent(1)
	tr.InsertWithParent(3)
	tr.InsertWithParent(5)
	tr.InsertWithParent(7)
	tr.InsertWithParent(9)

	tests := []struct {
		name string
		got  int
		want int
	}{
		{"case1", tr.Left.Parent.Value, 4},
		{"case1", tr.Right.Right.Right.Parent.Value, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("Tree.InsertWithParent() = %v, want %v", tt.got, tt.want)
			}
		})
	}
}
