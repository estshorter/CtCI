package p08

import (
	"reflect"
	"testing"
)

func Test_showDuplicated(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 2, 3, 4, 5, 6, 2, 2, 2, 4, 5, 5, 32000, 0, 32000, 2999, 2999, 3000}}, []int{2999, 5, 2, 4, 32000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := showDuplicated(tt.args.a); !orderInsensitiveEqual(got, tt.want) {
				t.Errorf("countDuplicated() = %v, want %v", got, tt.want)
			}
		})
	}
}

func orderInsensitiveEqual(ss1, ss2 []int) bool {
	if len(ss1) != len(ss2) {
		return false
	}
	counter1 := make(map[int]int)
	counter2 := make(map[int]int)
	for _, s1 := range ss1 {
		counter1[s1]++
	}
	for _, s2 := range ss2 {
		counter2[s2]++
	}
	return reflect.DeepEqual(counter1, counter2)
}
