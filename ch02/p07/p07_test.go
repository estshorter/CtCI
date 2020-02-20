package p07

import (
	"ctci/container/sllist"
	"reflect"
	"testing"
)

func Test_detectCommonNode(t *testing.T) {
	l1 := sllist.InitBySlice([]int{1, 2, 3, 4, 5, 6, 7})
	l2 := sllist.InitBySlice([]int{1})
	e := l1.GetElement(5)
	for i := 0; i < 2; i++ {
		eNext := e.Next
		l2.PushBackElement(e)
		e = eNext
	}

	type args struct {
		l1 *sllist.List
		l2 *sllist.List
	}
	tests := []struct {
		name string
		args args
		want []*sllist.Element
	}{
		{"case1",
			args{l1, l2},
			[]*sllist.Element{l1.GetElement(5), l1.GetElement(6)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCommonNode(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCommonNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
