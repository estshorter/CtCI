package p12

import (
	"reflect"
	"testing"
)

func Test_placeQueen(t *testing.T) {
	tests := []struct {
		name string
		want [][gridsize]int
	}{
		{"1", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := placeQueen(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("placeQueen() = %v, want %v", got, tt.want)
			}
		})
	}
}
