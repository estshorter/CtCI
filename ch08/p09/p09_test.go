package p09

import (
	"reflect"
	"testing"
)

func Test_generateParens(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{3}, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParens(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParens() = %v, want %v", got, tt.want)
			}
		})
	}
}
