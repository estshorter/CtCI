package p07

import (
	"reflect"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	type args struct {
		mat              [][]int
		positiveRotation bool
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"2x2p",
			args{[][]int{{1, 2}, {3, 4}}, true},
			[][]int{{3, 1}, {4, 2}},
		},
		{"2x2n",
			args{[][]int{{1, 2}, {3, 4}}, false},
			[][]int{{2, 4}, {1, 3}},
		},
		{"5x5p",
			args{[][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}, true},
			[][]int{{21, 16, 11, 6, 1}, {22, 17, 12, 7, 2}, {23, 18, 13, 8, 3}, {24, 19, 14, 9, 4}, {25, 20, 15, 10, 5}},
		},
		{"5x5n",
			args{[][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}, false},
			[][]int{{5, 10, 15, 20, 25}, {4, 9, 14, 19, 24}, {3, 8, 13, 18, 23}, {2, 7, 12, 17, 22}, {1, 6, 11, 16, 21}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotateMatrix(tt.args.mat, tt.args.positiveRotation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
