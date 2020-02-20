package p08

import (
	"ctci/container/sllist"
	"reflect"
	"testing"
)

func Test_findBeginning(t *testing.T) {
	l1 := sllist.InitBySlice([]int{1, 2, 3, 4, 5})
	l1.PushBackLoop(6, 2)
	l2 := sllist.InitBySlice([]int{1, 2, 3, 4, 5, 6, 7})
	type args struct {
		l *sllist.List
	}
	tests := []struct {
		name string
		args args
		want *sllist.Element
	}{
		{"case1", args{l1}, l1.GetElement(2)},
		{"case2", args{l2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBeginning(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findBeginning() = %v, want %v", got, tt.want)
			}
		})
	}
}
