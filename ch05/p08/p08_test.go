package p08

import (
	"reflect"
	"testing"
)

func TestDrawLine(t *testing.T) {
	sc := make([]byte, 8)
	sc2 := make([]byte, 8)

	type args struct {
		screen []byte
		width  int
		x1     int
		x2     int
		y      int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"case1", args{sc, 32, 1, 30, 1}, []byte{0,0,0,0,127,255,255,254}},
		{"case2", args{sc2, 32, 2, 7, 0}, []byte{63,0,0,0,0,0,0,0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DrawLine(tt.args.screen, tt.args.width, tt.args.x1, tt.args.x2, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DrawLine() = %x, want %x", got, tt.want)
			}
		})
	}
}
