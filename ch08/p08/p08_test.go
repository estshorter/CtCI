package p08

import (
	"reflect"
	"testing"
)

func Test_buildRuneCount(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want map[rune]int
	}{
		{"1", args{"abc"}, map[rune]int{'a': 1, 'b': 1, 'c': 1}},
		{"2", args{"aabbcc"}, map[rune]int{'a': 2, 'b': 2, 'c': 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildRuneCount(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildRuneCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permutation(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"abc"}, []string{"abc", "acb", "bac", "bca", "cba", "cab"}},
		{"2", args{"aaaaa"}, []string{"aaaaa"}},
		{"3", args{"aabb"}, []string{"aabb", "abba", "abab", "baab", "baba", "bbaa"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutation(tt.args.input); !equal(got, tt.want) {
				t.Errorf("permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(ss1, ss2 []string) bool {
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
