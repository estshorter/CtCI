package p02

import (
	"reflect"
	"testing"
)

func orderInsensitiveEqual(ss1, ss2 []string) bool {
	if len(ss1) != len(ss2) {
		return false
	}
	counter1 := make(map[string]int)
	counter2 := make(map[string]int)
	for _, s1 := range ss1 {
		counter1[s1]++
	}
	for _, s2 := range ss2 {
		counter2[s2]++
	}
	return reflect.DeepEqual(counter1, counter2)
}

func Test_anagramGroupingSort(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{"oobarf", "case", "hoge", "acse", "oghe", "fuga", "foobar"}},
			[]string{"case", "acse", "hoge", "oghe", "fuga", "oobarf", "foobar"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anagramGroupingSort(tt.args.a); !orderInsensitiveEqual(got, tt.want) {
				t.Errorf("anagramGroupingSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
