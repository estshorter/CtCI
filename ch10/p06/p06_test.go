package p06

import (
	"reflect"
	"testing"
)

func Test_count(t *testing.T) {
	a := make([]int, 128-12)
	for i := 12; i < 128; i++ {
		a[i-12] = i
	}
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{a}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := count(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("count() = %v, want %v", got, tt.want)
			}
		})
	}
}
