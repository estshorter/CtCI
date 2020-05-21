package p04

import (
	"reflect"
	"testing"
)

func Test_subset(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{"a", "b", "c"}}, []string{"", "a", "b", "ab", "c", "ac", "bc", "abc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subset(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subset() = %v, want %v", got, tt.want)
			}
		})
	}
}
