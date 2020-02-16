package p05

import (
	"ctci/container/list"
	"fmt"
	"reflect"
	"testing"
)

func equal(l1 *list.List, l2 *list.List) bool {
	if l1.Len() != l2.Len() {
		fmt.Printf("wrong length: got %v, want %v\n", l1.Len(), l2.Len())
		return false
	}

	v1 := make([]int, l1.Len())
	v2 := make([]int, l2.Len())

	e1 := l1.Front()
	e2 := l2.Front()

	equal := true
	for i := 0; i < l1.Len(); i++ {
		v1[i] = e1.Value
		v2[i] = e2.Value

		if e1.Value != e2.Value {
			equal = false
		}
		e1 = e1.Next
		e2 = e2.Next
	}
	fmt.Printf("got %v, want %v\n", v1, v2)
	return equal
}

func Test_listToInt(t *testing.T) {
	l1 := list.NewWithSlice([]int{0, 3, 5})
	l2 := list.NewWithSlice([]int{4, 3, 2, 1})
	type args struct {
		l *list.List
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{l1}, 530},
		{"case2", args{l2}, 1234},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listToInt(tt.args.l); got != tt.want {
				t.Errorf("listToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intToList(t *testing.T) {
	l1 := list.NewWithSlice([]int{0, 3, 5})
	l2 := list.NewWithSlice([]int{4, 3, 2, 1})
	l3 := list.NewWithSlice([]int{0})
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want *list.List
	}{
		{"case1", args{530}, l1},
		{"case2", args{1234}, l2},
		{"case3", args{0}, l3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToList(tt.args.i); !equal(got, tt.want) {
				t.Errorf("intToList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	l1 := list.NewWithSlice([]int{0, 3, 5})
	l2 := list.NewWithSlice([]int{4, 3, 2, 1})
	l3 := list.NewWithSlice([]int{4, 6, 7, 1})

	type args struct {
		l1 *list.List
		l2 *list.List
	}
	tests := []struct {
		name string
		args args
		want *list.List
	}{
		{"case1", args{l1, l2}, l3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.l1, tt.args.l2); !equal(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listToIntReverse(t *testing.T) {
	l1 := list.NewWithSlice([]int{5, 3, 0})
	l2 := list.NewWithSlice([]int{1, 2, 3, 4})
	type args struct {
		l *list.List
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{l1}, 530},
		{"case2", args{l2}, 1234},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listToIntReverse(tt.args.l); got != tt.want {
				t.Errorf("listToIntReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intToListReverse(t *testing.T) {
	l1 := list.NewWithSlice([]int{5, 3, 0})
	l2 := list.NewWithSlice([]int{1, 2, 3, 4})
	l3 := list.NewWithSlice([]int{0})
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want *list.List
	}{
		{"case1", args{530}, l1},
		{"case2", args{1234}, l2},
		{"case3", args{0}, l3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToListReverse(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intToListReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumReverse(t *testing.T) {
	l1 := list.NewWithSlice([]int{5, 3, 0})
	l2 := list.NewWithSlice([]int{1, 2, 3, 4})
	l3 := list.NewWithSlice([]int{1, 7, 6, 4})

	type args struct {
		l1 *list.List
		l2 *list.List
	}
	tests := []struct {
		name string
		args args
		want *list.List
	}{
		{"case1", args{l1, l2}, l3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumReverse(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
