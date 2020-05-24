package p10

import (
	"reflect"
	"testing"
)

func Test_fill(t *testing.T) {
	row := 5
	col := 2
	img1 := make([][]Color, row)
	for i := 0; i < row; i++ {
		img1[i] = make([]Color, col)
	}
	img1[1][1] = white
	img1[0][1] = white
	img1[1][0] = white
	img2 := make([][]Color, row)
	for i := 0; i < row; i++ {
		img2[i] = make([]Color, col)
	}

	type args struct {
		img       [][]Color
		row       int
		col       int
		fillColor Color
		p         Point
	}
	tests := []struct {
		name string
		args args
		want [][]Color
	}{
		{"1", args{img1, 5, 2, red, Point{4, 0}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fill(tt.args.img, tt.args.row, tt.args.col, tt.args.fillColor, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fill() = %v, want %v", got, tt.want)
			}
		})
	}
}
