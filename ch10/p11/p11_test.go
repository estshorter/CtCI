package p11

import (
	"reflect"
	"testing"
)

func Test_sortPeakTrough(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{5, 3, 1, 2, 3}}, []int{2, 1, 3, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortPeakTrough(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPeakTrough() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortPeakTrough2(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{5, 3, 1, 2, 3}}, []int{3, 5, 1, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortPeakTrough2(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPeakTrough2() = %v, want %v", got, tt.want)
			}
		})
	}
}
