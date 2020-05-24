package p07

import (
	"reflect"
	"testing"
)

func Test_permutation(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"bc"}, []string{"bc", "cb"}},
		{"2", args{"abc"}, []string{"abc", "bac", "bca", "acb", "cab", "cba"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutation(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
