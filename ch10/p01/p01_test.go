package p01

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	a := [10]int{1, 4, 6, 10}
	a2 := [10]int{1, 4, 6, 10}
	ret := [10]int{1, 1, 2, 4, 5, 6, 10}
	ret2 := [10]int{1, 1, 2, 4, 5, 6, 10, 12, 14, 20}
	type args struct {
		a  []int
		b  []int
		na int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{a[:], []int{1, 2, 5}, 4}, ret[:]},
		{"2", args{a2[:], []int{1, 2, 5, 12, 14, 20}, 4}, ret2[:]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.a, tt.args.b, tt.args.na); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge2(t *testing.T) {
	a := [10]int{1, 4, 6, 10}
	a2 := [10]int{1, 4, 6, 10}
	ret := [10]int{1, 1, 2, 4, 5, 6, 10}
	ret2 := [10]int{1, 1, 2, 4, 5, 6, 10, 12, 14, 20}

	type args struct {
		a  []int
		b  []int
		na int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{a[:], []int{1, 2, 5}, 4}, ret[:]},
		{"2", args{a2[:], []int{1, 2, 5, 12, 14, 20}, 4}, ret2[:]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge2(tt.args.a, tt.args.b, tt.args.na); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge2() = %v, want %v", got, tt.want)
			}
		})
	}
}
