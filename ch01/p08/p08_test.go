package p08

import (
	"reflect"
	"testing"
)

func Test_zeronize(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"2x2p",
			args{[][]int{{0, 2, 0}, {4, 5, 6}}},
			[][]int{{0, 0, 0}, {0, 5, 0}},
		},
		{"2x2n",
			args{[][]int{{1, 2}, {3, 4}, {5, 0}}},
			[][]int{{1, 0}, {3, 0}, {0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zeronize(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zeronize() = %v, want %v", got, tt.want)
			}
		})
	}
}
