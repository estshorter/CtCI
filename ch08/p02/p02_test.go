package p02

import (
	"reflect"
	"testing"
)

func Test_getPathNaive(t *testing.T) {
	const row = 5
	const col = 4
	maze := make([][]bool, row)
	for r := 0; r < row; r++ {
		maze[r] = make([]bool, col)
		for c := 0; c < col; c++ {
			maze[r][c] = true
		}
	}
	maze[1][0] = false
	maze[1][1] = false
	want1 := []Point{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2}, {4, 3}}
	type args struct {
		maze [][]bool
	}
	tests := []struct {
		name string
		args args
		want []Point
	}{
		{"1", args{maze}, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPath(tt.args.maze); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
